# Music2Noteblocks.

This Repository is an AI-Powered Audio and Song analysis tool.
The core of this program is a custom-built multi-layered Machine Learning model for generating Machine-readable encodings of the input song. 
This encoding will be used to design automated implementation instructions. These instructions shall be for building a virtual orchestra of the input song complient with the second-most popular metaverse.

# Music2Noteblocks: The Discryption I used when I wasn't maxing out on buzzwords in an attempt to boost my salary.
I going to build a Minecraft music machine myself? No! Imma make a machine that makes from the music the machine that makes the minecraft music machine.

# Impetus
Some time ago I signed up for one of those famous "200 players" Minecraft civilization event. 
I am not particularly skilled at minecraft: I am not abreast of the latest updates. My redstone knowledge is 5 years out-of-date. When building, my inventory fills up, and work is delayed. 
My PVP strategy involves dumping a bucket lava upon my opponent. It works well enough if I score a critical hit, but that's not guranteed.
Even if I off my fellow player, all their equipment promptly gets burned. What's the point of that, then? 

Some folk, though organization, dedication, and inspiration, are elected leaders. Foreign Policy! I would very much appreciate such a post, if I could get one.
I never managed such to land the top-job, though.

What's something with sufficient Pizzazz to get my otherwise-mediocre skills some attention?

# Noteblock Music
Minecraft has Noteblocks. They can be set to different pitches. Stringing multiple notblocks together, one can make some impressive songs (https://youtu.be/KyElxl_j4Wc?si=Z8qAx-1wWWHHZZwc&t=11).
It's impressive.
If I could lay claim to such an achievement, people will be impressed. I would have gained prestige in the Minecraft Civilization Event.

# What, am I going to do that by myself?
The problem with making note-block songs is that they're a pain in the bum. 
А. I have a good enough ear for Myst's Piano Rocket. Automatically translating a song that I'm hearing in real-time into note notation, that I cannot do. Triply so if said song has, say, 3 instruments, which will in turn require 3 types of noteblocks.
Б. Noteblocks don't have a UI. You have to Right-click them a certain number of times to get the right pitch. Go one over? Now you have to Right-click until it resets. 
When does it re-set? You've already mis-counted, so who knows? (Now that I've had time to think about it, I suppose I could just mine & re-place the Noteblock. This does not abrogate the other points! )
В. It requires counting the spaces in between notes. Also in real-time. ~~(Sheet music would mitigate some of these problems, but where would I get such a sheet music?)~~
Г. Building the machine is boring. 

# Solution: Mindflayer
Once, when I was looking for other Open-Source projects to work on, I stumbled upon Mineflayer (https://github.com/PrismarineJS/mineflayer).
Perfect! I bot myself, have the bot build the music contraption, and then log back in and take all the credit! 
Hopefully, building a Noteblock song is tedious and monotonous enough that this bot is indistinguishable to a rather bored player doing the work.
But how be the block bot born? This question is what this program seeks to answer. It takes in an audio-format song, dissects it into its constituent notes, which are then programmed as "Place this noteblock here" instructions in the bot.

# Chosen Audio format: .ogg
.ogg is an odd audio format. Nonetheless, it is the format chosen:
1. Minecraft music is .oog. Seeing as this is a program by Minecraft, for Minecraft, it makes sence for that to be the training data. Furthermore, Minecraft is a popular enough game such ~~that sheet music is likely readily available.~~ (Oops! See what happens when you assume?)
2. Wikipedia (mostly) uses .ogg. If I need to "jazz up": my training data, I doubt it'll be hard to find sheet music for "Ode To Joy" (https://en.wikipedia.org/wiki/Ode_to_Joy)

# Notes on training Data.
Minecraft's time is based off of ticks. 20 Ticks make a second. 10 ticks is 200 seconds. If the end result is a noteblock note for each tick, the end result of my AI is a 200-second vector, with each entry denoting the note played during that particular tick.
200 is a large vector size. It's big enough. Thus, this Gizmo will also read in only 10 seconds of song at at time. This prevents my Gizmo from getting too complacent  about song size. 

# Language: Go. Why Not?
Conveniently, around the time I was writing out the .readme for this program, I heard of a certain talk at one of the many networking events I go to. (I Go to a lot of networking events because I have no job)
The talk is about "MCP and AG-UI protocols along with LangChain to build an AI powered workflow Python."
Un-conveniently, I have no idea what any of those mean. I Like Python. GoLang is also a nice language, though.
