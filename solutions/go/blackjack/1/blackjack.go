package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        	return 11
        case "two":
        	return 2
        case "three":
        	return 3
        case "four":
        	return 4
        case "five":
        	return 5
        case "six":
        	return 6
        case "seven":
        	return 7
        case "eight":
        	return 8
        case "nine":
        	return 9
        case "king", "queen", "jack", "ten":
        	return 10
        default:
        	return 0	
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardone := ParseCard(card1)
    cardtwo := ParseCard(card2)
    hands := cardone + cardtwo
    dealer := ParseCard(dealerCard)
    if hands == 21 && dealer < 10 {
        return "W"
    } else if hands > 16 && hands < 21 {
        return "S"
    } else if (hands > 11 && hands <= 16) && dealer >= 7 {
        return "H"
    } else if hands == 22 {
        return "P"
    } else if hands == 21 && (dealer == 10 || dealer == 11) {
      	return "S"  
    } else if (hands > 11 && hands <= 16) {
        return "S"
    } else {
        return "H"
    }
}
