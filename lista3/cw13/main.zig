const std = @import("std");

var rng: std.rand.Xoroshiro128 = undefined;

fn foo(x: f64) f64 {
    return @sqrt(1.0 - x * x);
}

fn monteCarloPriv(a: f64, b: f64, n: u16, bound: f64, f: *const fn (f64) f64) f64 {
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

    return res;
}

fn monteCarloStratifiedPriv(a: f64, b: f64, n: u16, bound: f64, f: *const fn (f64) f64) f64 {
    const mid = (a + b) / 2.0;

    const resLeft = monteCarloPriv(a, mid, n / 2, bound, f);
    const resRight = monteCarloPriv(mid, b, n / 2, bound, f);

    return resLeft + resRight;
}

fn monteCarloStratified(n: u16) f64 {
    return monteCarloStratifiedPriv(0, 1, n, 1.0, foo);
}

fn monteCarloSimple(n: u16) f64 {
    return monteCarloPriv(0, 1, n, 1.0, foo);
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

    const numOfRuns = 10_000;

    const expResult = std.math.pi / 4.0;
    var resStrat: f64 = 0.0;
    var resSimple: f64 = 0.0;
    var resAntithetic: f64 = 0.0;
    for (0..numOfRuns) |_| {
        resStrat += @abs(monteCarloStratified(n) - expResult);
        resSimple += @abs(monteCarloSimple(n) - expResult);
        resAntithetic += @abs(monteCarloAntithetic(n) - expResult);
    }
    resStrat /= @floatFromInt(numOfRuns);
    resSimple /= @floatFromInt(numOfRuns);
    resAntithetic /= @floatFromInt(numOfRuns);

    std.debug.print("Avg Result err for simple: {}\n", .{resSimple});
    std.debug.print("Avg Result err for stratified: {}\n", .{resStrat});
    std.debug.print("Avg Result err for antithetic: {}\n", .{resAntithetic});
}
