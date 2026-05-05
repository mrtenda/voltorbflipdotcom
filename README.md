## [VoltorbFlip.com](https://voltorbflip.com)

![Voltorb Flip Screenshot](https://cloud.githubusercontent.com/assets/1281326/21961444/b72a3f76-dabe-11e6-8cdc-91fac816452c.png)

Voltorb Flip is a puzzle minigame found in the Goldenrod Game Corner and Celadon Game Corner in Pokemon Heart Gold and Soul Silver for the Nintendo DS. It is a game of both skill and chance that plays like a cross between Picross and Minesweeper. It replaces the traditional slot machines found in past Game Corners, and appears in all languages of the games, except for the Japanese release, which instead has slot machines. For an in-depth explanation of the rules of Voltorb Flip, see [Bulbapedia's article on Voltorb Flip](http://bulbapedia.bulbagarden.net/wiki/Voltorb_Flip).

VoltorbFlip.com is a website that takes Voltorb Flip puzzles as input and automatically solves them as efficiently as possible. Since Voltorb Flip has an element of chance to it though, this solver cannot successfully solve every single puzzle 100% of the time. However, it does tell the user the best moves they can make in any puzzle.

### How It Works

Given a board as input, which includes the row/column totals as well as any flipped cards' values, the solver follows a simple algorithm:

1. Run a set of heuristic algorithms against the inputted board which deduce any possible information they can about values the unflipped cards on the board. One example of one of these heuristic algorithms is "if the number of remaining hidden Voltorbs plus the number of remaining hidden points in a row totals to exactly the number of unflipped cards, then it is known that all unflipped cards in this row must have a value of 1". There are several such heuristic algorithms that this program uses. I came up with some of these heuristic algorithms myself, but some of the more inventive ones came from [this Voltorb Flip guide](http://www.dragonflycave.com/johto/voltorb-flip).
2. If there are still any unflipped 2 or 3 cards remaining, check if there are any cards that are known not to be Voltorbs on the board. If so, ask the user to flip those cards, since they are completely safe to flip.
3. If there are still any unflipped 2 or 3 cards remaining, create a list of all possible valid solutions to the current board. This is done using a breadth-first search.
4. Using this list, find the unflipped card on the board that is the safest guess. The safest card to guess is the one that is a Voltorb the least amount of times in all the possible solutions.
5. Prompt the user to flip the safest card found in the previous step.
6. Repeat until the user either finds all of the 2 and 3 cards (which means they won) or the user finds a Voltorb (which means they lost).

### How to Deploy (AWS Serverless)

#### Prerequisites
- AWS CLI configured with credentials.
- Node.js & npm.
- Docker (mandatory for the build process).

#### Deployment
```powershell
cd infra
npx cdk deploy
```

#### Environment Configuration
The stack is configured to deploy to **us-east-1** (N. Virginia) to support CloudFront SSL certificates and Route 53 lookups. To change the target account or region, update the `env` block in `infra/bin/infra.ts`.

### Local Development (Legacy Docker)
You can still build and run the project locally using the provided Dockerfile. Note that the backend has been updated to be Lambda-compatible, so you may need to use a Lambda emulator (like AWS SAM or LocalStack) for full local testing of the API.
