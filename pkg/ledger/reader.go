package ledger

import (
	"bufio"
	"os"
)

type Transaction struct {
	Date        string
	Description string
	Credit      string
	Debit       string
	Amount      float64
}

// ReadFile reads a file and returns its content as a slice of strings.
func ReadFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ParseTransactions parses a slice of strings into a slice of Transaction objects.
// func ParseTransactions(lines []string) ([]Transaction, error) {
// 	var transactions []Transaction

// 	for _, line := range lines {
// 		// Skip empty lines and comments
// 		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
// 			continue
// 		}

// 		// Parse each line into Transaction fields
// 		fields := strings.Fields(line)
// 		if len(fields) < 3 {
// 			return nil, fmt.Errorf("invalid line format: %s", line)
// 		}

// 		amount, err := parseAmount(fields[len(fields)-1])
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to parse amount: %w", err)
// 		}

// 		transaction := Transaction{
// 			Date:        fields[0],
// 			Description: strings.Join(fields[2:], " "),
// 			Credit:      fields[1],
// 			Debit:       "",
// 			Amount:      amount,
// 		}

// 		transactions = append(transactions, transaction)
// 	}

// 	return transactions, nil
// }

// func parseAmount(amountStr string) (float64, error) {
// 	amountStr = strings.TrimPrefix(amountStr, "€")
// 	amount, err := strconv.ParseFloat(amountStr, 64)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return amount, nil
// }
