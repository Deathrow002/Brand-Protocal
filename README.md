# Brand-Protocal

1. Boss's Baby Revenge

The program takes a user-provided string as input, representing a sequence of events:
- "S" signifies a water gun shot fired by the kids at Boss Baby's house.
- "R" represents Boss Baby retaliating with a water gun shot.

Based on the analysis, the program provides one of three outputs:
- "Good Boy" - Boss Baby followed the "fair revenge" protocol.
- "Bad Boy" - Boss Baby initiated shooting or retaliated excessively.

Example: 

        Input Sequence Event: SSRRR
        Input: SSRRR
        Output: Bad Boy  # Boss Baby retaliated more than he was shot at.

####################################################################################

2. Superman's Chicken Rescue

The program prompts you for two inputs:
- Superman Roof: The height of Superman's roof (integer).
- Chicken Positions: Comma-separated integers representing the positions of each chicken on the building.

The program performs the following steps:
- Conversion: It attempts to convert the comma-separated chicken positions into a list of integers. If the input contains non-numeric characters, it displays an error message and returns 0.
- Calculation: It iterates through each chicken's position and calculates Superman's maximum reach for that location (considering roof height).
- Rescue Count: It checks if Superman's reach is sufficient to save subsequent chickens. If the reach is enough, it increments a counter for the number of potentially rescued chickens from that position.
- Maximum Rescue: It compares the rescue count from each position and keeps track of the highest number of chickens Superman can save.

Example:

        Input Superman Roof: 7
        Enter Superman Chickens Position 'comma-separated integers': 2, 4, 8, 10
        Maximum Chickens Rescued: 3

####################################################################################

3. Transaction Broadcasting and Monitoring Document


This document describes how to use and how it work.

Running the Script
- go run TransactionBroadcastingMonitoring.py

Transaction Status:
- If successful (status code 200), the script will display the transaction hash and repeatedly check the transaction status using the check API endpoint (https://mock-node-wgqbnxruha-as.a.run.app/check/).
- The script will continue checking until the transaction status is confirmed, failed, or "DNE" (does not exist).
- For successful transactions, the script will display the confirmation message.
- For failed transactions, the script will display the error message.
- For non-existent transactions, the script will display an error message.

- Example:

        $ go run TransactionBroadcastingMonitoring.go
        
        Input Symbol: BTC
        Input Price: 12000
        Data sent successfully!
        Transaction Hash: 57a4071f55d0ee110a91dd7e1a1b8168692696dcbaa49e2a919962ba04fc6893
        Transaction Status: PENDING
        Transaction Status: PENDING
        Transaction Status: PENDING
        Transaction Status: PENDING
        Transaction Status: PENDING
        Transaction Status: CONFIRMED