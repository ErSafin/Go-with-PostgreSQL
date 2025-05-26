package main

import (
       "fmt"
	   "time"
	   "os"
       "database/sql"
       "log"
      _ "github.com/lib/pq" 
	  
   )

   type ticket struct {
       Airport      string 
       Salary       string 
	   Data_ticket  string 
	   Name         string    
    }


   func main() {
	   
    TicketQuery := "SELECT ticket.airport,ticket.salary,ticket.data_ticket,passenger.name FROM ticket JOIN passenger ON ticket.passengerid=passenger.id AND "
	  	   
      var pm int 
	  var sql_dt string
	  
      fmt.Println("Menu: \n1.Find at Data. \n2.Find at Person. \n3.Find at Airport")
      fmt.Println("---------------------------------------")
      fmt.Println("Choice menu (Enter symbol): ")
      fmt.Fscan(os.Stdin, &pm)
      switch pm {
        case 1:
          fmt.Println("Enter data: ")
          fmt.Fscan(os.Stdin, &sql_dt)   
		  sql_dt1, err := time.Parse("2006-01-02", sql_dt) // будет тип Time
			if err!=nil{
				panic("no DATA")
			}
			sql_dt=sql_dt1.Format("2006-01-02") // будет тип string
			TicketQuery = TicketQuery + "ticket.date="+"'"+sql_dt+"'"
        case 2: 
	      fmt.Println("Enter Person: ")
          fmt.Fscan(os.Stdin, &sql_dt) 
		  TicketQuery = TicketQuery + "passenger.name="+"'"+sql_dt+"'"
        case 3:
          fmt.Println("Enter Airport: ")
          fmt.Fscan(os.Stdin, &sql_dt) 
		  TicketQuery = TicketQuery + "ticket.airport="+"'"+sql_dt+"'"
        default:
          panic("noable menu")  
        
       }                  
       
     /* sql_dt - данные для поиска и выборки из БД  
     
	        SELECT ticket.airport,ticket.salary,ticket.data_ticket,passenger.name
            FROM ticket 
            JOIN passenger ON ticket.passengerid=passenger.id AND ticket.date=sql_dt
	  */
       connStr := "user=postgres dbname=ticket sslmode=disable"
       dbase, err := sql.Open("postgres", connStr)
       if err != nil {
           log.Fatal(err)
       }
	   
       defer dbase.Close()
       
	   rows, err := dbase.Query(TicketQuery)
	   if err != nil {
           log.Fatal(err)
       }
       defer rows.Close()	
	   
	  tks := []ticket{}
	  
       for rows.Next() {
           t:= ticket{}
		   err := rows.Scan(&t.airport, &t.salary,&t.data_ticket)
		   if err != nil {
               log.Fatal(err)
			   continue
           }
	       tks = append(tks, t)
       }
	   fmt.Println(tks)

	   for _, t:=range tks {
		   fmt.Println(t.airport, t.salary, t.data_ticket, t.name)
	   }

    }





