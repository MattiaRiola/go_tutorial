# Guards
There are 𝑁 prisoners in jail cells numbered from 0 to 𝑁−1 and 𝐺 guards. The guards ensure the prisoners do not escape and can only take charge of adjacent segments of jail cells. Each jail cell is assigned to exactly one guard.

Each prisoner 𝑖 has a certain intelligence 𝑆𝑖, and if the guard is watching him has 𝑘 people to watch over, his escaping possibility will be 𝑘𝑆𝑖. You should assign the guards so that the total escaping possibility over all the prisoners is minimized.

## input

The first line of input contains 2 integers 𝑁 and 𝐺 (1≤8000≤𝑁,1≤3000≤𝐺).

The next 𝑁 lines of input will contain one integer each, 𝑆𝑖 (1≤𝑆𝑖≤10^9).

## output

The output should contain one line with one integer, the minimum total escaping possibility.

## Example

| input | input |
| ----- | ----- |
| 6     | 3     |
| 11    |       |
| 11    |       |
| 11    |       |
| 24    |       |
| 26    |       |
| 100   |       |


| output |
| ------ |
| 299    |

