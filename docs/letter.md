# Love letter to my past professional self

Topics: computer programming, software engineering, automated software testing.

Dear me in 2013,

You're doing great as a junior engineer. You've mastered many things, and your pride equals your contribution to the corporation.

As you know, that one big software responsibility needs even more work to solve persistent deployment problems. Your manager has suggested the idea of making code to test the code, and I think they are right. The attention to detail you've learned that's needed to manually test and share code changes successfully is good, as is your creativity in thinking of features to speed up making changes, but you could be more effective by adding a variety of automation to your verification.

Fortunately since leaving the company to do entrepreneurship I've implemented some kinds of automated software testing along the way that I think would be helpful to you. In this letter I'll share some of those ideas.

I define automated tests as more code to test the code. Usually these aren't built into the regular program executable file or included with the installation of it. They are extra programs you run while developing the software that might show defects, and they are often designed for the project.

Like how a sports team has a practiced set of game tactics, the author of a software project pulls from a playbook of testing techniques. Automated testing only provides some testing plays, so don't assume your project is done when test programs don't report problems. Automatic tests should subtract from problems you are responsible for fixing, but then keep being smart about what else you do to know you can say it's done. A compiler doesn't print every problem in source code, and similarly automated tests won't show every problem in your computer application.

Within the automated subset of testing there are infinite ideas of what to try. The following sections detail a few I've explored for my local network web server application, but within the information, expertise, or thinking you have available to you there are likely many more plays. Stay creative.

## Debugging with unit tests

You know that isolating the code mistake that causes a problem sometimes requires lots of speculation within many functions, with a feeling maybe similar to walking through a swamp. Investigating with a hypothesis means keeping many assumptions in mind, and inspecting these assumptions is needed to continue onto the next hypothesis when one fails to explain the problem.

You've heard of unit testing and understand the idea of how to do it, but verifying every functionality in a project with unit tests seems like a waste of time. All work hours are already allocated, and the code is done when it does what it's supposed to do which is obvious. These tests also add unnecessary work when they often need to be changed to keep up with regular changes.

This mindset of agility and careful manual verification instead of unit testing, which might be valid for a style of project, could be like throwing the champagne out with the cork. If you are doing speculative debugging and are thinking through many assumptions unit tests can be a helpful technique to prove individual assumptions.

My project's application addressing system could have been the cause of a problem, but a small program to test the related functions (a unit test) showed that it was implemented as expected, so therefore the problem was something else.

```
package rules

import "testing"

func TestAddressToIndex(t *testing.T) {
	addresses := []Address{{0, 0}, {7, 7}, {3, 5}, {6, 7}, {6, 0}, {7, 0}, {0, 7}}
	indices := []AddressIndex{0, 63, 43, 62, 6, 7, 56}
	for i, addr := range addresses {
		if addr.Index() != indices[i] {
			t.Fatal("address", addr, "doesn't match index", indices[i])
		}
		if indices[i].Address() != addr {
			t.Fatal("index", indices[i], "doesn't match address", addr)
		}
	}
}
```

This is a unit test in the Go programming language for my Go code. Go build tools implement a unit testing specification that the above code is using (they're run by the ```go test``` command), but making a regular program to do testing would be easy. In C just make a separate program that links in the symbols you want to test, and perhaps a script could build and run all test programs and print a message if one returns a non-zero code to indicate the test failed.

The above test proved that my problem symptoms weren't related to those functions, and I was able to find the mistake elsewhere. This test will always be available to prove that this code is correct even if its implementation is changed.

Your unit tests may grow into a productive automatic testing garden that provides a harvest of avoided mistakes. Usually these test programs are small and reasonable to consider done after some thinking; you shouldn't have to do another test to test the test.

(The source code and build script for the Go programming language build tools are where I saw a good example of unit testing, where a folder of many individual test programs are executed each time that project is built.)

## Case gardens

Sometimes a unit of functionality in a project normally accepts an enormous variety of inputs and produces as many outputs, so satisfactory test coverage feels almost impossible.

In my project these functions are the most complicated, have to be perfect, need to perform quickly, and are expected to be changed during important improvements. A very useful test of these functions for me has been ones that take a set of cases which define the expected results given an input, where it's easy to add new cases when I think of a way it should be working.

Here's an example case and case file data structure and an excerpt from the test. JSON in Go is easy, so my test cases are encoded as JSON files that are read by the unit test.

```
type MovesCaseJSON struct {
	Name         string            `json:"case"`
	Active       rules.Orientation `json:"active"`
	PreviousMove rules.Move        `json:"prev"`
	State        rules.State       `json:"state"`
	Position     []testPiece       `json:"pos"`
	Moves        []rules.MoveSet   `json:"moves"`
}

type MovesCategoryFileJSON struct {
	Cases []MovesCaseJSON `json:"cases"`
}
```

The above structures looks like the following when saved in a file:

```
...
},
{
    "case": "Bad Check",
    "active": 0,
    "prev": {
        "f": {
            "f": 6,
            "r": 4
        },
        "t": {
            "f": 5,
            "r": 2
        }
    },
    "state": 0,
    "pos": [
        {
            "addr": {
                "f": 0,
                "r": 0
            },
            "k": 18,
            "o": 0,
            "m": false,
            "s": {
                "f": 0,
                "r": 0
            }
        },
        {
        ...
```

And the following is the test code excerpt:

```
func TestMoves(t *testing.T) {
	for _, tc := range loadAllMovesCases() {
	...
		// the complicated function being tested
		moves, state := board.Moves(tc.Active, tc.PreviousMove)
		if state != tc.State {
			t.Fatal(tc.Name, ":", "expected state", tc.State, "got", state)
		}
		...
```

This test currently verifies 38 cases for me and is very helpful when a problem is found during other testing; I add a new case that defines how the output should have been, then when the test passes I know I likely fixed that mistake and didn't add a new obvious mistake that would be caught by the other cases.

When a test will have many input cases it can be helpful to make a graphical user interface for quick case creation. This is another reason I used JSON encoding, since web browser JavaScript easily converts from data objects constructed by interface interactions into a JSON encoding that can then be sent to the web host program to be saved into a file. The GUI is very worth going out of your way to do if defining the test data in text is tedious, because easy entry means many more test cases which could be a big win for the quality of your project.

## Bingo hall

A test with a fun name I made up is the "bingo hall test" which simulates a bingo marathon in a hall with a hundred players (a metaphor for a hundred users of the software). By simulating parallel users I can find mistakes that cause crashes, and after it goes for a day I can assume that tens of users hopefully wouldn't easily experience a crash in the parts of the program that were tested.

This test uses another test program I wrote to benchmark the time of network requests. The below script implements the test by restarting that ```maxplayers``` benchmark every hour:

```
#!/usr/bin/env bash
i=1
while true
do
	echo hour $i
	./maxplayers -count 100 -length 3600 -host $1 -debug > bingo.log
	if [ $? -ne 0 ]; then
		echo maxplayers error
		exit 1
	fi
	((i=i+1))
done
```

When ```maxplayers``` encounters an unexpected response to its random but valid network requests it exits with code 1 and leaves the log file of every interaction it did. The host it's communicating with always prints information about what went wrong, including a stack trace and specific details which I carefully implemented to print, and using ```grep``` I can then isolate out of ```bingo.log``` the associated ```maxplayers``` order of actions that caused the crash. Usually I can translate this printed information into a test case which lets me quickly isolate and fix the mistake and then restart the bingo hall test.

After fixing tens of mistakes found by the bingo hall test I feel a sense of confidence in program correctness, but it's important to remember that automated testing isn't the only kind of testing you must do. Just because the program doesn't crash doesn't mean it will always work as expected, even with many runtime checks for invalid program states.

Manual testing isn't less important because of added automation, but it does feel good to fix mistakes found automatically that you could have learned of by a hard way that affected your users.

## Love your project

For my project the automated testing ideas above are just a start of what the complete effort could be. I'm still thinking of new tests, and there's plenty of formal methodologies described out there that could be a source of inspiration. Open source projects are also good to study for ideas.

Little is more important than being the person most involved in the project you are responsible for. When you have free moments tinker with it even more, and try to be its biggest fan.

Anyway, it's 2020 and computers still aren't as good as they could be. I know you feel the same way, so stay focused. Don't be so uptight though.

Sincerely,

Matt in 2020