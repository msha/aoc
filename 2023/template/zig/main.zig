const std = @import("std");

pub fn main() void {
    const hand = readFile("test.txt");

    for (hand) |card| {
        std.debug.print("{}\n", .{card});
    }

    // Your code goes here
}

fn readFile(filename: []const u8) ![][:0]const u8 {
    const file = try std.fs.cwd().openFile(filename, .{});
    defer file.close();

    var buffer = try file.readToEndAlloc(std.heap.page_allocator, std.math.maxInt(usize));
    defer std.heap.page_allocator.free(buffer);

    const lines = std.mem.tokenize(buffer, "\n");
    const result = try std.mem.allocSentinel(std.heap.page_allocator(), [:0]u8, lines.count, 0);

    var i: usize = 0;

    while (lines.next()) |line| {
        result[i] = line;
        i += 1;
    }

    return result;
}
