const std = @import("std");

var rng: std.rand.Xoroshiro128 = undefined;

fn foo(x: f64) f64 {
    return @sqrt(1.0 - x * x);
}

fn monteCarloPriv(a: f64, b: f64, n: u16, bound: f64, f: *const fn (f64) f64) [2]f64 {
    var hit: u64 = 0;
    var all: u64 = 0;

    for (0..n) |_| {
        const x = a + (b - a) * std.rand.float(rng.random(), f64);
        const y = std.rand.float(rng.random(), f64) * bound;
        all += 1;
        if (y <= f(x)) {
            hit += 1;
        }
    }

    var res: f64 = @floatFromInt(hit);
    res /= @floatFromInt(all);
    res *= (b - a) * bound;

    // TODO: Is this correct?

    var variance = res * (1.0 - res);
    variance /= @floatFromInt(all);

    return .{ res, variance };
}

fn monteCarloStratifiedPriv(a: f64, b: f64, n: u16, bound: f64, eps: f64, maxDepth: u16, f: *const fn (f64) f64) [2]f64 {
    if (n < 4) {
        return .{ 0.0, 0.0 };
    }
    const mid = (a + b) / 2.0;
    const newN = n / 2;

    const res1 = monteCarloPriv(a, mid, newN, bound, f);
    const res2 = monteCarloPriv(mid, b, newN, bound, f);

    const value = (res1[0] + res2[0]);
    const variance = (res1[1] + res2[1]) / 2.0;

    if (maxDepth == 0 or @abs(res1[1] - res2[1]) < eps or n < 4) {
        return .{ value, variance };
    }

    const res3 = monteCarloStratifiedPriv(a, mid, newN, bound, eps / 2.0, maxDepth - 1, f);
    const res4 = monteCarloStratifiedPriv(mid, b, newN, bound, eps / 2.0, maxDepth - 1, f);

    return .{ res3[0] + res4[0], res3[1] + res4[1] };
}

fn monteCarloStratified(n: u16, eps: f64, maxDepth: u16) f64 {
    return monteCarloStratifiedPriv(0, 1, n, 1.0, eps, maxDepth, foo)[0];
}

fn monteCarloSimple(n: u16) f64 {
    return monteCarloPriv(0, 1, n, 1.0, foo)[0];
}

fn monteCarloAntitheticPriv(a: f64, b: f64, n: u16, bound: f64, f: *const fn (f64) f64) f64 {
    var hit: u64 = 0;
    var all: u64 = 0;

    for (0..n) |_| {
        const rx = std.rand.float(rng.random(), f64);
        const ry = std.rand.float(rng.random(), f64);
        const x = a + (b - a) * rx;
        const y = ry * bound;
        all += 1;
        if (y <= f(x)) {
            hit += 1;
        }

        const x2 = a + (b - a) * (1.0 - rx);
        const y2 = (1 - ry) * bound;
        all += 1;
        if (y2 <= f(x2)) {
            hit += 1;
        }
    }

    var res: f64 = @floatFromInt(hit);
    res /= @floatFromInt(all);
    res *= (b - a) * bound;

    return res;
}

fn monteCarloAntithetic(n: u16) f64 {
    return monteCarloAntitheticPriv(0, 1, n, 1.0, foo);
}

pub fn main() !void {
    rng = std.rand.Xoroshiro128.init(seed: {
        var s: u64 = 0;
        std.posix.getrandom(std.mem.asBytes(&s)) catch |err| {
            std.debug.panic("Failed to get random seed: {}\n", .{err});
        };
        break :seed s;
    });
    const n = 1000;
    const eps = 0.0001;
    const maxDepth = 10;

    const resStrat = monteCarloStratified(n, eps, maxDepth);
    const resSimple = monteCarloSimple(n);
    const resAntithetic = monteCarloAntithetic(n);

    std.debug.print("Result for simple: {}\n", .{resSimple});
    std.debug.print("Result for stratified: {}\n", .{resStrat});
    std.debug.print("Result for antithetic: {}\n", .{resAntithetic});
}
