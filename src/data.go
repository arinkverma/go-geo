package main

import(
	"bufio"
	"strings"
	"os"
	"log"
	"fmt"
)

const CityFile = "cities15000.txt"

type DataContext struct{
	redisCtx *RedisContext
}

func (ctx DataContext) InitData(){
	file, err := os.Open(CityFile)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	// Purge the existing data
	conn := ctx.redisCtx.Get()
	defer conn.Close()
	_, errFlush := conn.Do("flushall")
	if errFlush != nil{
		log.Fatal("Can't purge older data", errFlush)
	}
	log.Printf("Purged existing data")

	count := 0
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	count += 1
    	ctx.addRecord(scanner.Text())
    }
    log.Printf("Total %d records have been updated.", count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}

func (ctx DataContext) addRecord(data string){
	words := strings.Split(data, "\t")
	geonameid := words[0]
	name := words[1]
	latitude := words[4]
	longitude := words[5]
	countryCode := words[8]
	value := fmt.Sprintf("%s:%s:%s", geonameid,name,countryCode)
	ctx.redisCtx.GeoAdd(latitude, longitude, value)
}
