# https://www.acmicpc.net/problem/2178
from collections import deque

N, M = map(int, input().split())
miro = []

for _ in range(N):
    miro.append(list(map(int, input())))

queue = deque()

visit = [[0]*M for _ in range(N)]
queue.append([0,0,1])
visit[0][0] = 1

while queue:
    x, y, c = queue.popleft()
    visit[x][y] = 1
    if x == N-1 and y == M-1:
        print(c)
        break
    for dx, dy in [(0, 1), (1,0), (0, -1), (-1, 0)]:
        nx, ny = x+dx, y+dy
        if nx < 0 or ny < 0 or nx >= N or ny >= M:
            continue
        if miro[nx][ny] == 1 and visit[nx][ny] == 0:
            queue.append([nx, ny, c+1])
    




# 우하우하 이동
