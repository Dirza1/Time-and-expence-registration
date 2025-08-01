/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

// overviewDatesCmd represents the overviewDates command
var OverviewDatesType string
var OverviewDatesFirstDate string
var OverviewDatesSecondDate string

var overviewDatesCmd = &cobra.Command{
	Use:   "overviewDates",
	Short: "Gives an overview of either or both of the databases between two specified dates.",
	Long: `This command returns an overview of one or both of the databaases entries between the two specified dates.
	One flag sets the datase to be querries and the other two flags specify the dates to querry.`,
	Run: func(cmd *cobra.Command, args []string) {
		firstDate := utils.TimeParse(OverviewDatesFirstDate)
		querry := utils.DatabaseConnection()

		secondDate := utils.TimeParse(OverviewDatesSecondDate)

		money := database.OverviewTransactionsDateParams{}
		time := database.OverviewTimeDatesParams{}
		switch OverviewDatesType {
		case "Finance":
			money = database.OverviewTransactionsDateParams{
				DateTransaction:   firstDate,
				DateTransaction_2: secondDate,
			}
			fmt.Println("Overview op financial databse:")
			fmt.Println(querry.OverviewTransactionsDate(context.Background(), money))
		case "Time":
			time = database.OverviewTimeDatesParams{
				DateActivity:   firstDate,
				DateActivity_2: secondDate,
			}
			fmt.Println("Overview op timeregistration databse:")
			fmt.Println(querry.OverviewTimeDates(context.Background(), time))
		case "All":
			money = database.OverviewTransactionsDateParams{
				DateTransaction:   firstDate,
				DateTransaction_2: secondDate,
			}
			fmt.Println("Overview op financial databse:")
			fmt.Println(querry.OverviewTransactionsDate(context.Background(), money))
			time = database.OverviewTimeDatesParams{
				DateActivity:   firstDate,
				DateActivity_2: secondDate,
			}
			fmt.Println("Overview op timeregistration databse:")
			fmt.Println(querry.OverviewTimeDates(context.Background(), time))
		default:
			fmt.Println("Incorrect use of the type flag. Use Finance, Time or All. Pay mind to the capitalation.")
		}
	},
}

func init() {
	rootCmd.AddCommand(overviewDatesCmd)

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesType, "type", "t", "all", "Flag to specify the database to querry. Use Finance, Time or All after the flag")
	err := overviewDatesCmd.MarkFlagRequired("type")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesFirstDate, "First-Date", "f", "", "Flag to specify the first date to querry. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")
	err = overviewDatesCmd.MarkFlagRequired("month")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	overviewDatesCmd.Flags().StringVarP(&OverviewDatesSecondDate, "Second-Date", "s", "", "Flag to specify the second date to querry. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")
	err = overviewDatesCmd.MarkFlagRequired("year")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overviewDatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overviewDatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
