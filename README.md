# JsonLab

CLI tool to deal with all things JSON.

## Install

I'd recommend you to build this from source. It's easy if you have Go
installed and `$GOCODE/bin` in your `$PATH`:

```
go get github.com/aybabtme/jsonlab
```

## Usage

For now, it only does string encoding/decoding. It works on UTF8 strings,
so you can do random stuff like:

```
$ pzalgo < sample.txt | jsonlab strings escape | tee -a /dev/stderr | jsonlab strings unescape

"To invoke t̀he hiͧve-mind ͧrepͬres͖ēntͭi̱n̰ͫ̀g ̵cháȯś̮.̞\nͧ͘Ḭň̛v̫o̟k̗ͯi̿nͮ͘g͆҉͓ ͎th̡͉ͨe͑ ̢̫ͩf͉eͪ҉e̛͐lͯi̥̅n͏͕ǵ ̴̼o̗͢f̉͝ ̷͍̇c̅͏h͈ͧ͠ȁo̵̼ͩs̶̙̓.̮̇͡\ǹ̫́Wͭ͝i͓t̛̹h̩̔̕ ̝̀͝o̡̠͂ű̱͠t̓͏̥ ̪̽ȏ̶ͅr̍҉͈d̿̀ͅe̋҉͈r̹ͭ̕.̢̻ͤ\n̷͓ͨT̲ͨ͠h̫͛͝ĕ̝̀ ͚̐͟N̵̲͌e͇̚͟z͚͆̀p͌͏̟e̸̡̥̣͗̊r̛̦̣̐̅́d̵̩̩̒̃̀i͔̤̊͋͘͝a̵̛͔͂̓ͅn̵̛̤̹͊ͤ ̖̗̅̏̀͢h̏̊͏̻̹͡ḯ̸̸̞̟̑v̢̨͚̺͛͒e̷̥͍͗̅͡-̵͓̗̐ͭ͞m̷̳̜ͧ͆́i̶̥̖ͥ͐̕n̶̡̰̜͛ͯḑ̸̖͚̍̄ ͎̥͐͌̀̕ŏ̵͉̗ͭ͝f̫͙ͩ̔͟͟ ̧̲͇ͯ̈́͟c̙͕̐̾͡͞h̶̨͇̠́̈́a̐́҉̬͈́o̱ͤ̆̕͜ͅs̢͇̫ͮͩ͘.̠͉̾͊͟͝ ̴͈͇̈́͌͞Ż̶̢̼̻ͫå̛͉̭ͤ͡ļ͍͔ͮ͂́g̤̞͆̽̀͢o͖͕ͫ̎̕͠.̬͈̆ͫ́͡\n̜͉ͪ̎́͝H̃̾͢͏͚̩ḛ̭̃ͥ͘͜ ̢̽ͩ͏͇̰w͑ͪ̀҉̳͕h̷͎͍̓̇͜oͧ̾҉̴̖̝ ̸̶̰̹̑̏W̴̫̬̅͒͞ą̸̱̦ͩ̿į̴̘͇ͩ͆t̵̠̪ͬ̄͞s̛͍̭͋̆͝ ̡̧̥̙͂̿B͗ͭ͏̦̘̀é̚҉̶̮ͅh̬̦͊̂͘͘ĩ͚̘̾̕͞n̴̦͇ͭͫ͢d̸͉̹͂̆͜ ̨͖͚ͪ͌͢T̋̀͏͔̼̀ḩ̸̩̳͒̽eͧͯ̕͏̼ͅ ͉̗̎̉̕͡W̨̤̞ͨͬ́aͦͧ͏͖̼͢l̸̨̰̜̿̿l̴̢̼̱͑̇.̛̫̞̈̈͞\ň͐҉̛͇̰Z̡̺̦͂̔͟Ą͉̼́ͦ́L̴̵͔̲͛ͬGͯͦ͏̛͙̹O̢͚̦̐ͫ͜!̷̲̳ͭͦ͘\n̨̰̳̋͋̀"

To invoke t̀he hiͧve-mind ͧrepͬres͖ēntͭi̱n̰ͫ̀g ̵cháȯś̮.̞
 ͧ͘Ḭň̛v̫o̟k̗ͯi̿nͮ͘g͆҉͓ ͎th̡͉ͨe͑ ̢̫ͩf͉eͪ҉e̛͐lͯi̥̅n͏͕ǵ ̴̼o̗͢f̉͝ ̷͍̇c̅͏h͈ͧ͠ȁo̵̼ͩs̶̙̓.̮̇͡
 ̫̀́Wͭ͝i͓t̛̹h̩̔̕ ̝̀͝o̡̠͂ű̱͠t̓͏̥ ̪̽ȏ̶ͅr̍҉͈d̿̀ͅe̋҉͈r̹ͭ̕.̢̻ͤ
 ̷͓ͨT̲ͨ͠h̫͛͝ĕ̝̀ ͚̐͟N̵̲͌e͇̚͟z͚͆̀p͌͏̟e̸̡̥̣͗̊r̛̦̣̐̅́d̵̩̩̒̃̀i͔̤̊͋͘͝a̵̛͔͂̓ͅn̵̛̤̹͊ͤ ̖̗̅̏̀͢h̏̊͏̻̹͡ḯ̸̸̞̟̑v̢̨͚̺͛͒e̷̥͍͗̅͡-̵͓̗̐ͭ͞m̷̳̜ͧ͆́i̶̥̖ͥ͐̕n̶̡̰̜͛ͯḑ̸̖͚̍̄ ͎̥͐͌̀̕ŏ̵͉̗ͭ͝f̫͙ͩ̔͟͟ ̧̲͇ͯ̈́͟c̙͕̐̾͡͞h̶̨͇̠́̈́a̐́҉̬͈́o̱ͤ̆̕͜ͅs̢͇̫ͮͩ͘.̠͉̾͊͟͝ ̴͈͇̈́͌͞Ż̶̢̼̻ͫå̛͉̭ͤ͡ļ͍͔ͮ͂́g̤̞͆̽̀͢o͖͕ͫ̎̕͠.̬͈̆ͫ́͡
 ̜͉ͪ̎́͝H̃̾͢͏͚̩ḛ̭̃ͥ͘͜ ̢̽ͩ͏͇̰w͑ͪ̀҉̳͕h̷͎͍̓̇͜oͧ̾҉̴̖̝ ̸̶̰̹̑̏W̴̫̬̅͒͞ą̸̱̦ͩ̿į̴̘͇ͩ͆t̵̠̪ͬ̄͞s̛͍̭͋̆͝ ̡̧̥̙͂̿B͗ͭ͏̦̘̀é̚҉̶̮ͅh̬̦͊̂͘͘ĩ͚̘̾̕͞n̴̦͇ͭͫ͢d̸͉̹͂̆͜ ̨͖͚ͪ͌͢T̋̀͏͔̼̀ḩ̸̩̳͒̽eͧͯ̕͏̼ͅ ͉̗̎̉̕͡W̨̤̞ͨͬ́aͦͧ͏͖̼͢l̸̨̰̜̿̿l̴̢̼̱͑̇.̛̫̞̈̈͞
 ̌͐҉̛͇̰Z̡̺̦͂̔͟Ą͉̼́ͦ́L̴̵͔̲͛ͬGͯͦ͏̛͙̹O̢͚̦̐ͫ͜!̷̲̳ͭͦ͘
 ̨̰̳̋͋̀%
```
Very smart.

## Download

You can use the following links to a build service. I don't know the authors
or if the binaries should be trusted.

[Reflections on Trusting Trust](http://cm.bell-labs.com/who/ken/trust.html)
> Ken Thompson

---

* [Windows](http://gobuild.io/github.com/aybabtme/jsonlab/master/windows/amd64)
* [Linux](http://gobuild.io/github.com/aybabtme/jsonlab/master/linux/amd64)
* [OS X](http://gobuild.io/github.com/aybabtme/jsonlab/master/darwin/amd64)

The builds are provided by:

[![Gobuild Download](http://gobuild.io/badge/github.com/aybabtme/jsonlab/download.png)](http://gobuild.io/github.com/aybabtme/jsonlab)
