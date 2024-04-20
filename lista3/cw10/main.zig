const std = @import("std");

fn Array3D(comptime T: type) type {
    return struct {
        const Self = @This();
        allocator: std.mem.Allocator,
        x: usize,
        y: usize,
        z: usize,
        data: []T,

        pub fn init(allocator: std.mem.Allocator, x: usize, y: usize, z: usize) !Self {
            const data = try allocator.alloc(T, x * y * z);
            return .{
                .allocator = allocator,
                .x = x,
                .y = y,
                .z = z,
                .data = data,
            };
        }

        pub fn deinit(self: *Self) void {
            self.allocator.free(self.data);
        }

        pub fn get(self: *const Self, i: usize, j: usize, k: usize) T {
            return self.data[i * self.y * self.z + j * self.z + k];
        }

        pub fn set(self: *Self, i: usize, j: usize, k: usize, val: T) void {
            self.data[i * self.y * self.z + j * self.z + k] = val;
        }
    };
}
fn lcsForThree(comptime T: type, allocator: std.mem.Allocator, s1: []const T, s2: []const T, s3: []const T) ![]T {
    // Generate subproblems' array
    var array = try Array3D(usize).init(allocator, s1.len + 1, s2.len + 1, s3.len + 1);
    defer array.deinit();

    for (0..s1.len + 1) |i| {
        for (0..s2.len + 1) |j| {
            for (0..s3.len + 1) |k| {
                if (i == 0 or j == 0 or k == 0) {
                    array.set(i, j, k, 0);
                } else if (s1[i - 1] == s2[j - 1] and s2[j - 1] == s3[k - 1]) {
                    array.set(i, j, k, array.get(i - 1, j - 1, k - 1) + 1);
                } else {
                    const a = array.get(i - 1, j, k);
                    const b = array.get(i, j - 1, k);
                    const c = array.get(i, j, k - 1);
                    array.set(i, j, k, @max(@max(a, b), c));
                }
            }
        }
    }

    // Reconstruct the LCS
    var result = try allocator.alloc(T, array.get(s1.len, s2.len, s3.len));
    errdefer allocator.free(result);
    var i, var j, var k = .{ s1.len, s2.len, s3.len };
    var resultLen: usize = 0;
    while (i > 0 and j > 0 and k > 0) {
        const a = array.get(i - 1, j, k);
        const b = array.get(i, j - 1, k);
        const c = array.get(i, j, k - 1);

        const max = @max(@max(a, b), c);
        const addChar = max != array.get(i, j, k);
        if (a == max) {
            i -= 1;
            if (addChar) {
                result[resultLen] = s1[i];
                resultLen += 1;
            }
        } else if (b == max) {
            j -= 1;
            if (addChar) {
                result[resultLen] = s2[j];
                resultLen += 1;
            }
        } else if (c == max) {
            k -= 1;
            if (addChar) {
                result[resultLen] = s3[k];
                resultLen += 1;
            }
        }
    }
    std.mem.reverse(T, result);

    return result;
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const alc = gpa.allocator();
    defer {
        const status = gpa.deinit();
        if (status == .leak) {
            @panic("Memory leak detected.\n");
        }
    }
    const s1: []const u8 = "AGGT12T";
    const s2: []const u8 = "12TXTAY";
    const s3: []const u8 = "12XBATT";
    const result = try lcsForThree(u8, alc, s1, s2, s3);
    defer alc.free(result);
    std.debug.print("LCS: {s}\n", .{result});
    std.debug.print("LCS's length: {}\n", .{result.len});
}
