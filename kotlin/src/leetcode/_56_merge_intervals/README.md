# 56. Merge Intervals 

--- 
[← Back to Main](../../../../README.md)

## Problem
주어진 구간들을 병합하여 겹치지 않는 구간들의 배열을 반환합니다.

[LeetCode 링크](https://leetcode.com/problems/merge-intervals/)

## Approach

### 1. Greedy
- 시작점 기준 정렬 후 순차적으로 병합
- Time: O(n log n), Space: O(n)

## Key Points
- 정렬 후에는 인접한 구간만 확인하면 됨
- 겹침 조건: `end >= nextStart`
- 병합 시: `end = maxOf(end, nextEnd)`