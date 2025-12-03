#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <limits>
#include <algorithm> 
#include <iterator>
#include <stack>

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

long long solvePart1(const std::vector<std::string>& inputLines) {
    long long sum = 0;

    for (const auto& line: inputLines) {
        int max1 = -1, max2 = -1;
        int ix = 0;
        int posMax1 = 0;

        for (char c : line) {
            if (ix == line.size()-1) {
                break;
            }
            int num = c - '0';

            if (num > max1) {
                max1 = num;
                posMax1 = ix;
            }
            ix++;
        }

        for (int iy = posMax1+1; iy < line.size(); iy++) {
            int num = line[iy] - '0';

            if (num > max2) {
                max2 = num;
            }
        }

        sum += (max1*10 + max2);
    }
    return sum;
}

long long solvePart2(const std::vector<std::string>& inputLines) {
    long long sum = 0;
    for (const std::string& line : inputLines) {
        std::stack<char> stack;
        int toRemove = line.size() - 12;

        for (char c : line) {
            while(!stack.empty() && toRemove > 0 && stack.top() < c) {
                stack.pop();
                toRemove--;
            }
            stack.push(c);
        }
        while (stack.size() > 12) {
            stack.pop();
        }
        std::string out;

        while (!stack.empty()) {
            out += stack.top();
            stack.pop();
        }
        std::reverse(out.begin(), out.end());
        //std::cout << "out: " << std::stoll(out) << std::endl;
        sum += std::stoll(out);
    }
    return sum;
}