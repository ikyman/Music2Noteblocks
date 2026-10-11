var(
	DEFAULT_BOT_FOLDER string
)

struct BotSettings{
	floorBlocks []string
	deleteExistingFloor bool
	deleteableBlocks []string

	start_x int
	start_y int
	start_z int

	direction string //Temporarily a string, because I didn't get jazzed by any of the three AI-Generated Enum Values.
	width int

	errorBehavior string
}

func init(){
	DEFAULT_BOT_FOLDER = "./MineflayerBotWriter/MineflayerBots"
}

func WriteMineflayerBot(botName string){
	filePath := filepath.Join(defaultBotFolder, botName + ".js")

	err := os.MkdirAll(defaultBotFolder, 0755)
	if err != nil {
		panic(err)
	}

	botCreation := fmt.Sprintf(
	"""const mineflayer = require('mineflayer')
	const vec3 = require('vec3')
	
	if (process.argv.length < 4 || process.argv.length > 6) {
	  console.log('Usage : node %s.js <host> <port> [<name>] [<password>]')
	  process.exit(1)
	}
	
	const bot = mineflayer.createBot({
	  host: process.argv[2],
	  port: parseInt(process.argv[3]),
	  username: process.argv[4] ? process.argv[4] : '%s',
	  password: process.argv[5]
	})""");

	err = os.WriteFile(filePath, botCreation, 0644)
	if err != nil {
		panic(err)
	}

}