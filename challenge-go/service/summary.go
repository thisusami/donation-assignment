package service

import "fmt"

type SummaryService struct {
	Calculator CalculateInterface
}

func (s *SummaryService) Summarize() error {
	fmt.Println("\nperforming donations...")
	calculator, err := s.Calculator.Read()
	if err != nil {
		return fmt.Errorf("failed to read data: %w", err)
	}
	summary, err := calculator.Calculate()
	if err != nil {
		return fmt.Errorf("failed to calculate summary: %w", err)
	}
	fmt.Println("done.")
	fmt.Println()
	fmt.Printf("        total received: THB %10.2f\n", summary.Total)
	fmt.Printf("  successfully donated: THB %10.2f\n", summary.Successfully)
	fmt.Printf("       faulty donation: THB %10.2f\n", summary.Faulty)
	fmt.Println()
	fmt.Printf("    average per person: THB %10.2f\n", summary.Average)
	fmt.Printf("            top donors: ")
	for i, item := range summary.TopDonator {
		if i == 0 {
			fmt.Printf("%s\n", item)
		} else {
			fmt.Printf("                        %s\n", item)
		}
	}
	return nil
}
func NewSummary(calculator CalculateInterface) *SummaryService {
	return &SummaryService{
		Calculator: calculator,
	}
}
