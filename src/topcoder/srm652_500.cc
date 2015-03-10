#include <iostream>
#include <sstream>
#include <cstdio>
#include <cstdlib>
#include <cmath>
#include <cstring>
#include <cctype>
#include <string>
#include <vector>
#include <list>
#include <queue>
#include <deque>
#include <stack>
#include <map>
#include <set>
#include <algorithm>
#include <numeric>

using namespace std;

typedef long long Int;
typedef pair<int,int> PII;
typedef vector<int> VInt;

#define FOR(i, a ,b) for( i = (a); i < (b); ++i )
#define RFOR(i, a , b) for( i = (a) - 1; i >= (b); --i)
#define CLEAR(a, b) memset(a, b, sizeof(a))
#define SIZE(a) int((a).size())
#define ALL(a) (a).begin(),(a).end()
#define PB push_back
#define MP make_pair

#define MAX (1 << 10)
#define INF (1LL << 60)

Int R[MAX][MAX];
Int C[MAX];
vector<PII> E[MAX];
vector<PII> B[MAX];

class MaliciousPath {
public:
  long long minPath(int N, int K, vector <int> from, vector <int> to, vector <int> cost) {
    int i, j, k;
    FOR(i, 0, K + 1)
      FOR(j, 0, N)
        FOR(j, 0, N)
          R[i][j] = INF;

    FOR(i, 0, SIZE(from)) {
      int a = from[i];
      int b = to[i];
      int c = cost[i];
      E[a].PB(PII(b,c));
      B[b].PB(PII(a,c));
    }

    CLEAR(C,0);
    FOR(i, 0, K + 1){
      R[i][N - 1] = 0;
      priority_queue<pair<Int, int> > Q;
      Q.push(MP(-R[i][N - 1], N - 1));
      while (!Q.empty()) {
        int a = Q.top().second;
        Int r = -Q.top().first;
        Q.pop();

        if (r != R[i][a]) continue;

        FOR(j, 0, SIZE(B[a])) {
          int b = B[a][j].first;
          Int v = max(r + B[a][j].second, C[b]);
          if (v < R[i][b]) {
            R[i][b] = v;
            Q.push(MP(-v, b));
          }
        }
      }

      CLEAR(C, 0);
      FOR(j, 0, N)
        FOR(k, 0, SIZE(E[j]))
          C[j] = max(C[j], R[i][E[j][k].first] + E[j][k].second);
    }

    FOR(i, 0, N) {
      E[i].clear();
      B[i].clear();
    }

    return R[K][0] == INF ? -1 : R[K][0];
  }
};
