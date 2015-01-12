int count(int[] stones) {
int piles = sizeof(stones)/sizeof(*stones);
// loop every item see either all even or all odd
int index;
int even=stones[0]%2 == 0;
int sum = stones[0];
for (index=1; index!=piles; index++) {
  if (stones[index]%2 ^ even)
    return -1;
  else sum+=stones[index];
}
if (sum%piles!=0 && ((sum/piles)%2==0 ^ even))
return -1;


int target = sum / piles;
int steps=0;
for (index=0; index!=piles; index++) {
  steps += (stones[index] - target) / 2;
}
return steps/2;
}
