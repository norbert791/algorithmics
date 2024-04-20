const std = @import("std");

const GraphEdge = [2]u32;

const SimpleGraph = struct {
    arena: std.heap.ArenaAllocator,
    adjacencyList: []?[]u32,

    fn edgeLessThan(_: void, a: GraphEdge, b: GraphEdge) bool {
        return a[0] < b[0];
    }

    pub fn init(allocator: std.mem.Allocator, edges: []GraphEdge) !SimpleGraph {
        const duplEdges = try duplicateEdges(allocator, edges);
        defer allocator.free(duplEdges);
        std.sort.heap(
            GraphEdge,
            duplEdges,
            {},
            edgeLessThan,
        );

        var maxVertex: u32 = 0;
        for (0..duplEdges.len) |i| {
            if (duplEdges[i][0] > maxVertex) {
                maxVertex = duplEdges[i][0];
            }
            if (duplEdges[i][1] > maxVertex) {
                maxVertex = duplEdges[i][1];
            }
        }

        var arena = std.heap.ArenaAllocator.init(allocator);
        errdefer arena.deinit();
        var al = arena.allocator();
        var list = try al.alloc(?[]u32, maxVertex + 1);
        for (0..list.len) |i| {
            list[i] = null;
        }

        // Create adjacency lists
        var currentEdgeIndex: u32 = 0;
        for (0..duplEdges.len) |i| {
            if (duplEdges[currentEdgeIndex][0] != duplEdges[i][0]) {
                const adjList = try al.alloc(u32, i - currentEdgeIndex);
                for (currentEdgeIndex..i) |j| {
                    adjList[j - currentEdgeIndex] = duplEdges[j][1];
                }
                list[duplEdges[currentEdgeIndex][0]] = adjList;
                currentEdgeIndex = @intCast(i);
            }
        }
        // Append the last adjacency list
        const adjList = try al.alloc(u32, duplEdges.len - currentEdgeIndex);
        for (currentEdgeIndex..duplEdges.len) |j| {
            adjList[j - currentEdgeIndex] = duplEdges[j][1];
        }
        list[duplEdges[currentEdgeIndex][0]] = adjList;

        return SimpleGraph{
            .arena = arena,
            .adjacencyList = list,
        };
    }

    pub fn deinit(self: SimpleGraph) void {
        self.arena.deinit();
    }

    pub fn duplicateEdges(alloc: std.mem.Allocator, edges: []GraphEdge) ![]GraphEdge {
        var newEdges = try alloc.alloc(GraphEdge, edges.len * 2);
        errdefer alloc.free(newEdges);

        for (0..edges.len) |i| {
            newEdges[i] = edges[i];
            newEdges[i + edges.len] = .{ edges[i][1], edges[i][0] };
        }
        return newEdges;
    }
};

pub fn maxCut(allocator: std.mem.Allocator, graph: SimpleGraph) ![2][]u32 {
    var set1 = std.AutoArrayHashMap(u32, void).init(allocator);
    defer set1.deinit();
    var set2 = std.AutoArrayHashMap(u32, void).init(allocator);
    defer set2.deinit();

    for (0..graph.adjacencyList.len) |i| {
        const w: u32 = @intCast(i);
        if (graph.adjacencyList[w] == null) {
            try set1.put(w, {});
            continue;
        }
        var deltaSet1: u32 = 0;
        var deltaSet2: u32 = 0;
        for (graph.adjacencyList[w].?) |v| {
            if (set1.get(v) != null) {
                deltaSet1 += 1;
            } else if (set2.get(v) != null) {
                deltaSet2 += 1;
            }
        }
        if (deltaSet1 > deltaSet2) {
            try set2.put(w, {});
        } else {
            try set1.put(w, {});
        }
    }

    const res1 = blk: {
        const temp = set1.keys();
        const res = try allocator.alloc(u32, temp.len);
        errdefer allocator.free(res);
        std.mem.copyForwards(u32, res, temp);
        break :blk res;
    };
    errdefer allocator.free(res1);

    const res2 = blk: {
        const temp = set2.keys();
        const res = try allocator.alloc(u32, temp.len);
        errdefer allocator.free(res);
        std.mem.copyForwards(u32, res, temp);
        break :blk res;
    };
    errdefer allocator.free(res2);

    return .{ res1, res2 };
}
pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer {
        const status = gpa.deinit();
        if (status == .leak) {
            @panic("Memory leak detected.\n");
        }
    }
    var alloc = gpa.allocator();
    var edges = [_]GraphEdge{
        .{ 0, 1 },
        .{ 0, 2 },
        .{ 1, 2 },
        .{ 1, 3 },
        .{ 2, 3 },
        .{ 3, 4 },
    };
    const graph = try SimpleGraph.init(alloc, edges[0..]);
    defer graph.deinit();
    if (graph.adjacencyList.len != 5) {
        @panic("Graph adjacency list length is not 5\n");
    }
    const res = try maxCut(alloc, graph);

    const s1, const s2 = res;
    defer {
        alloc.free(s1);
        alloc.free(s2);
    }

    std.debug.print("graph: {any}\n", .{graph.adjacencyList});
    std.debug.print("first set: {any}\n", .{s1});
    std.debug.print("second set: {any}\n", .{s2});
}
