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
fn lcsForThree(comptime T: type, allocator: std.mem.Allocator, s1: []const T, s2: []const T, s3: []const T) !usize {
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
    return array.get(s1.len, s2.len, s3.len);
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
    const s1: []const u8 = "AGGT12";
    const s2: []const u8 = "12TXAY";
    const s3: []const u8 = "12XBA";
    const result = try lcsForThree(u8, alc, s1, s2, s3);
    std.debug.print("Length of LCS is: {}\n", .{result});
}
