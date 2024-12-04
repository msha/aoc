const std = @import("std");

pub fn main() !void {
    var allocator = std.heap.page_allocator;
    const input = try readFile(&allocator, "input.txt"); // Pass a reference to allocator

    defer allocator.free(input);

    var data = try parseInput(allocator, input);
    defer freeData(allocator, &data);

    var instructions = data.instructions;
    var fullMap = &data.map;
    var currentTiles = getCurrentTiles(allocator, fullMap);
    var steps: usize = 0;

    while (!allTilesAtZ(currentTiles)) {
        const currentInstruction = instructions[steps % instructions.len];
        currentTiles = getNewNodes(allocator, currentTiles, fullMap, currentInstruction);
        if (steps % 10000 == 0) {
            std.debug.print("{d}\n", .{steps});
        }
        steps += 1;
    }

    std.debug.print("{d}\n", .{steps});
}

fn readFile(allocator: *std.mem.Allocator, path: []const u8) ![]u8 {
    const file = try std.fs.cwd().openFile(path, .{});
    defer file.close();

    const len = try file.getEndPos();
    const buffer = try allocator.alloc(u8, len);

    try file.readAll(buffer);
    return buffer;
}

fn parseInput(allocator: *std.mem.Allocator, input: []u8) !ParsedData {
    var lines = std.mem.split(input, "\n");
    var iter = lines.iterator();

    const instructionsLine = try iter.next();
    if (instructionsLine == null) return error.MissingInstructions;

    var mapItems = try allocator.alloc(Node, lines.len - 1); // Allocating array for map nodes
    var index: usize = 0;

    while (true) {
        const line = iter.next() orelse break;
        if (line.len == 0) continue; // Skip empty lines

        var parts = std.mem.split(line, " ");
        const node = try parts.next().?; // Expecting node part

        const lrPart = try parts.next().?;
        const lrSplit = std.mem.split(lrPart[1 .. lrPart.len - 1], ",");
        const L = try lrSplit.next().?;
        const R = try lrSplit.next().?;

        mapItems[index] = Node{
            .node = try allocator.dupe(u8, node),
            .L = try allocator.dupe(u8, L),
            .R = try allocator.dupe(u8, R),
        };
        index += 1;
    }

    return ParsedData{
        .instructions = try allocator.dupe(u8, instructionsLine),
        .map = mapItems[0..index],
    };
}

fn getCurrentTiles(allocator: *std.mem.Allocator, fullMap: []Node) ![]Node {
    var currentTiles = try allocator.alloc(Node, fullMap.len);
    var count: usize = 0;

    for (fullMap) |node| {
        if (node.node.len >= 3 and node.node[2] == 'A') {
            currentTiles[count] = node;
            count += 1;
        }
    }

    return currentTiles[0..count];
}

fn getNewNodes(allocator: *std.mem.Allocator, currentTiles: []Node, fullMap: []Node, instruction: u8) ![]Node {
    var newNodes = try allocator.alloc(Node, fullMap.len);
    var count: usize = 0;

    for (currentTiles) |tile| {
        for (fullMap) |mapNode| {
            if (std.mem.eql(u8, mapNode.node, tile.getDirection(instruction))) {
                newNodes[count] = mapNode;
                count += 1;
            }
        }
    }

    return newNodes[0..count];
}

fn allTilesAtZ(tiles: []Node) bool {
    for (tiles) |tile| {
        if (tile.node.len < 3 or tile.node[2] != 'Z') return false;
    }
    return true;
}

fn freeData(allocator: *std.mem.Allocator, data: *ParsedData) void {
    allocator.free(data.instructions);
    for (data.map) |node| {
        allocator.free(node.node);
        allocator.free(node.L);
        allocator.free(node.R);
    }
    allocator.free(data.map);
}

const Node = struct {
    node: []u8,
    L: []u8,
    R: []u8,
};

const ParsedData = struct {
    instructions: []u8,
    map: []Node,
};
