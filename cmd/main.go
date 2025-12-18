package main

func main() {
	cfg := config{
	addr: ":8090",
	db: dbConfig{

	},	
	}
	api := application {
		config: cfg,
	}
}