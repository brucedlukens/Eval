### Status
[![Build Status](https://travis-ci.org/brucedlukens/Eval.svg?branch=master)](https://travis-ci.org/brucedlukens/Eval.svg?branch=master)

# Simple Coding Evaluation

This task should take no longer than 2 hours to complete. 

Please complete this task in the programming language of your choice.

You can submit your work in one of two forms:
* Fork this repository and commit your work to the fork, then give us the url to your fork in Github and the url to your passing tests in TravisCI.
* Of, if you're sensitive to having your work viewable on Github, you may submit a zip file of your local clone of this repo.  The repo must contain configuration to make its' test work with TravisCI.

Write a program that takes these two files as input:

* [names.txt](https://raw.githubusercontent.com/PerformLine/Eval/master/names.txt) - Contains a list of names, one name per line.
* [list.txt](https://raw.githubusercontent.com/PerformLine/Eval/master/list.txt) - Contains a list of words, some are names, one per line.

Then prints (to stdout) the **line numbers** and **instance** in *list.txt* where names from *names.txt* are found.

For example, if *list.txt* contains three instances of the name 'james' it would output something like

```james: 51, 1450, 14123```

This means that the name 'james' appears in *list.txt* on those lines.

We are looking for the most efficient solution taking into account I/O, CPU and Memory usage as it is possible you solution may be needed to run against much larger data set, specifically one that may not fit completely in memory.  

You may state your assumptions and feel free to add additional justification as well. Please be prepared to discuss it when you come on site.

Lastly, your code should include meaningful tests.  Make sure your tests complete and your repo is instrumented to use the free TravisCI service to execute tests automatically.
