package transactionstest

import (
	"time"

	"github.com/bitcoin-sv/spv-wallet/models"
	"github.com/bitcoin-sv/spv-wallet/models/response"
)

func NewTransactions() *response.PageModel[response.Transaction] {
	return &response.PageModel[response.Transaction]{
		Content: []*response.Transaction{
			{
				Model: response.Model{
					Metadata: map[string]any{
						"domain":          "john.doe.test.4chain.space",
						"example_key1":    "example_key10_val",
						"ip_address":      "127.0.0.01",
						"user_agent":      "node-fetch",
						"paymail_request": "HandleReceivedP2pTransaction",
						"reference_id":    "1c2dcc61-f48f-44f2-aba2-9a759a514d49",
						"p2p_tx_metadata": map[string]any{
							"pubkey": "3fa7af5b-4568-4873-86da-0aa442ca91dd",
							"sender": "john.doe@handcash.io",
						},
					},
				},
				ID:                   "2c250e21-c33a-41e3-a4e3-77c68b03244e",
				Hex:                  "283b1c6deb6d6263b3cec7a4701d46d3",
				XpubOutIDs:           []string{"4c9a0a0d-ea4f-4f03-b740-84438b3d210d"},
				BlockHash:            "47758f612c6bf5b454bcd642fe8031f6",
				BlockHeight:          512,
				Fee:                  1,
				NumberOfInputs:       2,
				NumberOfOutputs:      3,
				TotalValue:           311,
				OutputValue:          100,
				Status:               "MINED",
				TransactionDirection: "incoming",
			},
			{
				Model: response.Model{
					Metadata: map[string]any{
						"domain":          "john.doe.test.4chain.space",
						"example_key1":    "example_key11_val",
						"ip_address":      "127.0.0.01",
						"user_agent":      "node-fetch",
						"paymail_request": "HandleReceivedP2pTransaction",
						"reference_id":    "8a1b677b-c52f-4c87-bfb9-3c771ad4decc",
						"p2p_tx_metadata": map[string]any{
							"pubkey": "30e737f7-438f-4720-b0f5-8da9c16b9440",
							"sender": "john.doe@handcash.io",
						},
					},
				},
				ID:                   "06c5ed12-9f37-4aba-b018-2da341efabd3",
				Hex:                  "68dd12011887ae60a639b22d11fbdaa1",
				XpubOutIDs:           []string{"50ff06c2-bdb8-44da-a56b-4bff00b7b116"},
				BlockHash:            "f038fd250cab29a4ca86082b280b2e31",
				BlockHeight:          1024,
				Fee:                  1,
				NumberOfInputs:       2,
				NumberOfOutputs:      2,
				TotalValue:           10,
				OutputValue:          100,
				Status:               "MINED",
				TransactionDirection: "incoming",
			},
		},
		Page: response.PageDescription{
			Size:          2,
			Number:        2,
			TotalElements: 2,
			TotalPages:    1,
		},
	}
}

func NewTransaction() *response.Transaction {
	return &response.Transaction{
		Model: response.Model{
			UpdatedAt: Parse("2024-02-26T11:02:28.069911Z"),
			CreatedAt: Parse("2024-02-26T11:00:28.069911Z"),
			Metadata: map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			},
		},
		ID:                   "07a1d71e-a1c2-4bca-abd0-7f006ad8dbe8",
		Hex:                  "0100039604ede024c6b6eb70279fe504e4eff",
		XpubInIDs:            []string{"59ab2244-d575-4bc8-987c-50ab9462ac12"},
		XpubOutIDs:           []string{"85952324-c42b-418c-8fa4-8ae3ec253460"},
		BlockHash:            "c5105d50088a17a08075ca542293a288",
		BlockHeight:          833505,
		Fee:                  1,
		NumberOfInputs:       3,
		NumberOfOutputs:      2,
		DraftID:              "d33f936d-475f-4641-9ac7-e0743ad146d6",
		TotalValue:           51,
		OutputValue:          50,
		Outputs:              map[string]int64{"1a08d547-3a72-4940-bfb8-c8e6dd7b5dd9": -51, "32d5a09c-ae30-4689-aa3a-908113f8e7d0": 50},
		Status:               "MINED",
		TransactionDirection: "outgoing",
	}
}

func NewDraftTransaction() *response.DraftTransaction {
	return &response.DraftTransaction{
		Model: response.Model{
			CreatedAt: Parse("2024-10-17T05:51:37.469411Z"),
			UpdatedAt: Parse("2024-10-17T07:51:37.469425+02:00"),
			Metadata: map[string]interface{}{
				"receiver": "john.doe.test4@john.doe.test.4chain.space",
				"sender":   "john.doe.test4@john.doe.test.4chain.space",
			},
		},
		ID:        "30bdc7ea-5d89-43ac-809c-5be58ab63b13",
		Hex:       "4cf0bd870147ed060cf31c472063f3b0",
		XpubID:    "fd0ae0c4-6bb5-448b-a523-7d4bbfbeb50d",
		ExpiresAt: Parse("2024-10-17T05:51:54.14184Z"),
		Configuration: response.TransactionConfig{
			ChangeDestinations: []*response.Destination{
				{
					Model:         response.Model{},
					ID:            "02361709-01a2-4ac5-929a-ece8f0b3eac1",
					XpubID:        "0a949f7d-4452-4d9c-baff-2f753603a380",
					LockingScript: "3322294a-6d79-4f60-84c4-8967af8c09ef",
					Type:          "pubkeyhash",
					Chain:         1,
					Num:           3,
					Address:       "c0e6d33c-62b4-44f7-917d-65e848e598f1",
					DraftID:       "49a5fdcc-78dc-4ea1-8e70-31b642d0d9f2",
				},
			},
			ChangeSatoshis: 18,
			FeeUnit: &response.FeeUnit{
				Satoshis: 1,
				Bytes:    1000,
			},
			Inputs: []*response.TransactionInput{
				{
					Utxo: response.Utxo{
						Model: response.Model{
							CreatedAt: Parse("2024-10-07T13:47:22.948193Z"),
							UpdatedAt: Parse("2024-10-17T05:51:37.455194Z"),
						},
						UtxoPointer: response.UtxoPointer{
							TransactionID: "e4f1cd33-95ca-4d82-9e24-f81cbfc76f83",
							OutputIndex:   0,
						},
						ID:           "7af40e57-abf9-4bd7-814a-88a91b28cc68",
						XpubID:       "de53520f-e014-4cdc-b1e1-0d1ab6aa4eca",
						Satoshis:     20,
						ScriptPubKey: "6bfec5df-f397-4117-9a06-7fa14ab8c91b",
						Type:         "pubkeyhash",
						DraftID:      "efe9ca32-7c8f-4d5e-89f6-47a96b15e66d",
						ReservedAt:   Parse("2024-10-17T05:51:37.454383Z"),
						SpendingTxID: "",
					},
					Destination: response.Destination{
						Model: response.Model{
							Metadata: map[string]any{
								"domain":          "john.doe.4chain.space",
								"ip_address":      "127.0.0.1",
								"paymail_request": "CreateP2PDestinationResponse",
								"reference_id":    "1302be9141691ff69d12c12323cf64b4",
								"satoshis":        float64(20), // encoding/json pkg by default interpret numbers as float64 values.
								"user_agent":      "node-fetch",
							},
						},
					},
				},
			},
			Outputs: []*response.TransactionOutput{
				{
					OpReturn: &response.OpReturn{},
					PaymailP4: &response.PaymailP4{
						Alias:           "john.doe.test4",
						Domain:          "john.doe.4chain.space",
						FromPaymail:     "from@domain.com",
						ReceiveEndpoint: "https://john.doe.test.4chain.space:443/v1/bsvalias/beef/{alias}@{domain.tld}",
						ReferenceID:     "f1d8dfab-f6ae-4c6f-94a7-5054a1853362",
						ResolutionType:  "p2p",
					},
					Satoshis: 1,
					Scripts: []*response.ScriptOutput{
						{
							Address:    "fe1278f1-02ea-45b1-bb9c-7cc57fb2677c",
							Satoshis:   1,
							Script:     "604b50c9-ed08-4c98-aea5-85191c5eda78",
							ScriptType: "pubkeyhash",
						},
					},
					To: "c238254d-c684-4b5f-9c39-1ea3d2bc68b1",
				},
			},
		},
		Status:    "acb56775-db4f-4a29-8211-b29d4bd151f6",
		FinalTxID: "dd556b23-30ad-4518-a140-b27c3236e32b",
	}
}

func Parse(s string) time.Time {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}
	return t
}

func NewResponseError(code int) *models.ResponseError {
	return &models.ResponseError{
		Code:    models.UnknownErrorCode,
		Message: "something-went-wrong",
	}
}

func NewSPVError(code int) *models.SPVError {
	err := NewResponseError(code)
	return &models.SPVError{
		Code:       err.Code,
		Message:    err.Message,
		StatusCode: code,
	}
}
