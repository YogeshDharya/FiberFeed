package dbservice

import (
	"github.com/nedpals/supabase-go"
)

func main(){
	supabaseURL := 
	supabaseKey := 
	tableName := 	
	client := supabase.CreateClient(supabaseURL, supabaseKey)
	data, err := client.From()	
}