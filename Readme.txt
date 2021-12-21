- Cron Expression Parser to see it's expanded format
----------+
contents: |
----------+
-Golang implementation
-Go testing

Introduction to cron Expression :

  +---------------- minute (0 - 59)
  |
  |    +------------- hour (0 - 23)
  |    |
  |    |    +---------- day of month (1 - 31)
  |    |    |
  |    |    |    +------- month (1 - 12)
  |    |    |    |
  |    |    |    |    +---- day of week (1 - 7) (Sunday=1)
  |    |    |    |    |
  |    |    |    |    |          +---- command
  |    |    |    |    |          |
 */15  0   1,15  *   1-5    /usr/bin/find

cron expression  ->   "*/15 0  1,15 * 1-5 /usr/bin/find" denotes cron will run  Every 15 minutes, between 00:00 and 00:59, on day 1 and 15 of the month, Sunday through Thursday
so its output will be

minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find

--------------------------------------------------------------------------------------------------------------------------------------------------------------


--------------------------------+
-Environment Setup on Mac OSX   |
--------------------------------+
make sure home brew is already install if not then follow instruction in the Additional section
step1:

    $ brew install golang

step2: make sure your shell config have following env variable

    $ vi ~/.bashrc (default bash)   or  $vi ~/.zshrc(for zsh)

step3:
      #copy and paste add these exports in the bash file:

     export GOPATH=/usr/local/opt/go
     export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
     export GOROOT=/usr/local/opt/go/libexec
     export GO111MODULE=off

step4:
    save and close your  terminal


----------------------+
Project build and run |
----------------------+
step 1: Extract project folder from CronExpressionReader.zip
      $ unzip CronExpressionReader.zip

step 2: Open a terminal and change directory to the project directory ie CronExpressionReader folder

      $ cd CronExpressionReader

step 3: To build the project

      $ go build CronExpressionReader.go

step 4:Expression needs to passed in quotes

      $ ./CronExpressionReader "*/15 0 1,15 * 1-5 /usr/bin/find"

step 5: To run test

      $ go test ./CronExpression -v




Additional
----------------------------------------+
Prerequisite Setup  on mac-OSx          |
----------------------------------------+
HomeBrew Installation

Run following command for HomeBrew Setup from the terminal
  $ xcode-select --install
  $ curl -fsSL -o install.sh https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh
  $ /bin/bash install.sh

more details (installation  guide for homebrew)
Good read:https://www.digitalocean.com/community/tutorials/how-to-install-and-use-homebrew-on-macos

