package db

import(
	"bufio"
	"strings"
	"os"
	"log"
	"fmt"
)

type DataContext struct{
	RedisCtx *RedisContext
	count int32
	DataFile string
}

func (self DataContext) InitData(){
	file, errFile := os.Open(self.DataFile)
	if errFile != nil{
		log.Fatal(errFile)
	}
	defer file.Close()

	conn := self.RedisCtx.Get()
	defer conn.Close()
	_, errDb := conn.Do("FLUSHALL")
	if errDb != nil{
		log.Fatal("Can't purge older data", errDb)
	}
	log.Printf("Purged existing data")

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	self.addRecord(scanner.Text())
    	self.count += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    log.Printf("Total %d records have been updated.", self.count)
}

func (self DataContext) addRecord(data string){
	words := strings.Split(data, "\t")
	geonameid := words[0]
	name := words[1]
	latitude := words[4]
	longitude := words[5]
	adminCode := words[7]
	countryCode := words[8]
	value := fmt.Sprintf("%s:%s:%s", geonameid, name, countryCode)
	if strings.HasPrefix(adminCode, "PPL") {
		self.RedisCtx.GeoAdd(latitude, longitude, value)
		self.count += 1
	}
}

func MakeDataContext(redisCtx *RedisContext, dataFile string) *DataContext{
	return &DataContext{
		RedisCtx: redisCtx,
		DataFile: dataFile,
	}
}
