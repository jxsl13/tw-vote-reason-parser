# tw-logs parser for vote kick and spec vote reasons


This is a basic parser of Teeworlds log files that can be utilized to create a .csv file with all unique vote reasons that can be found in your server log files.

Your log files go into the `logs`folder and they should have the suffix o either `.txt` or `.log`. Any other file will be skipped.

You build this application by simply executing `go build .` in this folder and then you will get an executable that you can run.

Execute `./tw-logs` in the current folder and you will see a new file that is called `reasons.csv`
This is like an Microsoft Excel spreadsheet file that can be opened in Microsoft Excel or Numbers (macOS).
There are two extra columns next to the `reason` column. This spreadsheet has the purpos of classifying reasons and associate them with corresponding actions. Depending on what context those reasons were used in.
The first context is the *kick vote*. So depending on what reason was used for kickvoting someone, you define an action.
The second context is the *spec vote*. Depending on what reason is used for moving a player via a vote to the spectators, an automated system may decide based on these classifications to either stop, voteban, kick, ban the voting player or simply to ignore the vote and let players decide.
