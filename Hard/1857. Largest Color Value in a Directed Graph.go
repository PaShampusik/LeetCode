package main

/*
There is a directed graph of n colored nodes and m edges. The nodes are numbered from 0 to n - 1.

You are given a string colors where colors[i] is a lowercase English letter representing the color of the ith node in this graph (0-indexed). You are also given a 2D array edges where edges[j] = [aj, bj] indicates that there is a directed edge from node aj to node bj.

A valid path in the graph is a sequence of nodes x1 -> x2 -> x3 -> ... -> xk such that there is a directed edge from xi to xi+1 for every 1 <= i < k. The color value of the path is the number of nodes that are colored the most frequently occurring color along that path.

Return the largest color value of any valid path in the given graph, or -1 if the graph contains a cycle.

Example 1:
Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
Output: 3
Explanation: The path 0 -> 2 -> 3 -> 4 contains 3 nodes that are colored "a" (red in the above image).
Example 2:

Input: colors = "a", edges = [[0,0]]
Output: -1
Explanation: There is a cycle from 0 to 0.
*/
func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)

	// Строим граф
	graph := make([][]int, n)
	indegree := make([]int, n)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	// dp[i][c] = максимальное количество цвета c на ЛЮБОМ пути, заканчивающемся в узле i
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}

	// Топологическая сортировка + DP одновременно
	queue := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
			// ТОЛЬКО сейчас инициализируем цвет стартового узла!
			colorIndex := int(colors[i] - 'a')
			dp[i][colorIndex] = 1
		}
	}

	processed := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		processed++

		// Обрабатываем всех соседей
		for _, neighbor := range graph[curr] {
			neighborColor := int(colors[neighbor] - 'a')

			// Обновляем DP для соседа
			for color := 0; color < 26; color++ {
				if color == neighborColor {
					// Если цвет соседа совпадает с этим цветом, добавляем +1
					dp[neighbor][color] = max(dp[neighbor][color], dp[curr][color]+1)
				} else {
					// Иначе просто переносим количество
					dp[neighbor][color] = max(dp[neighbor][color], dp[curr][color])
				}
			}

			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Проверяем цикл
	if processed != n {
		return -1
	}

	// Находим максимум
	result := 0
	for i := 0; i < n; i++ {
		for color := 0; color < 26; color++ {
			result = max(result, dp[i][color])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
