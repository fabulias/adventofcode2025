#include <vector>
#include <utility>
#include <algorithm>
#include <string>
#include <iostream>
#include <fstream>

long long solvePart1(const std::vector<std::string>& input);
long long solvePart2(const std::vector<std::string>& input);

int main(int argc, char* argv[]) {
    if (argc < 2) {
        std::cerr << "Usage: " << argv[0] << " <input_file>" << std::endl;
        return 1;
    }

    std::string filename = argv[1];
    std::ifstream inputFile(filename);

    if (!inputFile.is_open()) {
        std::cerr << "Error opening file: " << filename << std::endl;
        return 1;
    }

    std::vector<std::string> inputLines;
    std::string line;
    while (std::getline(inputFile, line)) {
        inputLines.push_back(line);
    }

    inputFile.close();

    long long resultPart1 = solvePart1(inputLines);
    std::cout << "Part 1 Result: " << resultPart1 << std::endl;

    long long resultPart2 = solvePart2(inputLines);
    std::cout << "Part 2 Result: " << resultPart2 << std::endl;

    return 0;
}

long long solvePart1(const std::vector<std::string>& input_lines) {
    long long fresh_count = 0;

    std::vector<std::pair<long long, long long>> ranges;
    std::vector<long long> queries;

    bool parse_ranges = true;
    for (const auto& line: input_lines) {
        if (line.empty()) {parse_ranges = false; continue;}
        if (parse_ranges) {
            int pos = line.find('-');
            long long l = std::stoll(line.substr(0, pos));
            long long r = std::stoll(line.substr(pos+1));
            ranges.push_back(std::make_pair(l, r));
        } else {
            // check ingredient IDs
            queries.push_back(stoll(line));
        }
    }

    // merge ranges
    std::sort(ranges.begin(), ranges.end());
    std::vector<std::pair<long long, long long>> merged_ranges;
    for (auto [l,r] : ranges) {
        if (merged_ranges.empty() || l > merged_ranges.back().second + 1) {
            merged_ranges.push_back(std::make_pair(l,r));
        } else {
            merged_ranges.back().second = std::max(r, merged_ranges.back().second);
        }
    }


    // Binary search
    for (auto query_val : queries) {
        auto it = upper_bound(merged_ranges.begin(), merged_ranges.end(), std::make_pair(query_val, LLONG_MAX));
        if (it != merged_ranges.begin()) {
            --it;
            if (it->first <= query_val && query_val <= it->second) fresh_count++;
        }
    }
    return fresh_count;
}

long long solvePart2(const std::vector<std::string>& input_lines) {
        std::vector<std::pair<long long, long long>> ranges;

        bool parse_ranges = true;
        for (const auto& line: input_lines) {
            if (line.empty()) {parse_ranges = false; continue;}
            if (parse_ranges) {
                int pos = line.find('-');
                long long l = std::stoll(line.substr(0, pos));
                long long r = std::stoll(line.substr(pos+1));
                ranges.push_back(std::make_pair(l, r));
            }
        }

        // merge ranges
        std::sort(ranges.begin(), ranges.end());
        std::vector<std::pair<long long, long long>> merged_ranges;
        for (auto [l,r] : ranges) {
            if (merged_ranges.empty() || l > merged_ranges.back().second + 1) {
                merged_ranges.push_back(std::make_pair(l,r));
            } else {
                merged_ranges.back().second = std::max(r, merged_ranges.back().second);
            }
        }

        long long total = 0;
        for (auto [l,r] :merged_ranges) {
            total += (r - l + 1);
        }
        return total;
}