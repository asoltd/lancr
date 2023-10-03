#!/bin/bash

curl -v -X 'PUT' \
'http://localhost:4200/v1/heroes/test-stub' \
-H 'accept: application/json' \
-H 'content-type: application/json' \
-H 'X-Firebase-ID-Token: eyJhbGciOiJSUzI1NiIsImtpZCI6IjlhNTE5MDc0NmU5M2JhZTI0OWIyYWE3YzJhYTRlMzA2M2UzNDFlYzciLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiUGlvdHIgT3N0cm93c2tpIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0xweF8yZ1dQMm0wVlByLVhiX2JyY1FxTEJpVGV2UDZyb1doWFM3Z2YzbHpUYz1zOTYtYyIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS9hc29sdGQtZ3VpbGRzIiwiYXVkIjoiYXNvbHRkLWd1aWxkcyIsImF1dGhfdGltZSI6MTY5NjI4MTM2NCwidXNlcl9pZCI6InpHZUs0aUxrRE1oU2t4cXR5eENOUFBJenNuNzMiLCJzdWIiOiJ6R2VLNGlMa0RNaFNreHF0eXhDTlBQSXpzbjczIiwiaWF0IjoxNjk2MjgxMzY0LCJleHAiOjE2OTYyODQ5NjQsImVtYWlsIjoicGlvdHJvc3RyQGdvb2dsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJnb29nbGUuY29tIjpbIjExNDc1ODU0OTgyNTE0NjI3MTQ5MyJdLCJlbWFpbCI6WyJwaW90cm9zdHJAZ29vZ2xlLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6Imdvb2dsZS5jb20ifX0.yy_J_9Jdiu7duO7fszlu1DLMdzim7C5qXW-Dkwzx1ZyjmlqKbWAdTOiA1L5T0HmttmIctF1VztJ6shLfhg8wQCZpHwvlGMgHg8YNkM2ySiTYCpNOhGtN0nm0g6bx2P9foQT2b4swJFwQoHZ-Or1XOwuqKf0ROTXB0-XirHfoCcbfzrRJk_w_lCQyHyOzP2WMdUeXfVbfJis9Euw8oY3YrTjUe42Rt3FdWi388abbzEyOjWeGzGgsrjpwzX_zW8GQP0lx1fu1Q0NuK7nLMz9UrpzEdtbJmuZIOVmEvlbksZ2drsw2eB0x8KTKFFhg7XMGJkQQ1sZ9zO7oUTTC1paRbA' \
-d '{"id":"test-stub","hero":{"email": "firestore is not that great :("}}'

curl -v -X 'POST' \
'http://localhost:4200/v1/heroes' \
-H 'accept: application/json' \
-H 'content-type: application/json' \
-H 'X-Firebase-ID-Token: eyJhbGciOiJSUzI1NiIsImtpZCI6IjlhNTE5MDc0NmU5M2JhZTI0OWIyYWE3YzJhYTRlMzA2M2UzNDFlYzciLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiUGlvdHIgT3N0cm93c2tpIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0xweF8yZ1dQMm0wVlByLVhiX2JyY1FxTEJpVGV2UDZyb1doWFM3Z2YzbHpUYz1zOTYtYyIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS9hc29sdGQtZ3VpbGRzIiwiYXVkIjoiYXNvbHRkLWd1aWxkcyIsImF1dGhfdGltZSI6MTY5NjI4MTM2NCwidXNlcl9pZCI6InpHZUs0aUxrRE1oU2t4cXR5eENOUFBJenNuNzMiLCJzdWIiOiJ6R2VLNGlMa0RNaFNreHF0eXhDTlBQSXpzbjczIiwiaWF0IjoxNjk2MjgxMzY0LCJleHAiOjE2OTYyODQ5NjQsImVtYWlsIjoicGlvdHJvc3RyQGdvb2dsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJnb29nbGUuY29tIjpbIjExNDc1ODU0OTgyNTE0NjI3MTQ5MyJdLCJlbWFpbCI6WyJwaW90cm9zdHJAZ29vZ2xlLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6Imdvb2dsZS5jb20ifX0.yy_J_9Jdiu7duO7fszlu1DLMdzim7C5qXW-Dkwzx1ZyjmlqKbWAdTOiA1L5T0HmttmIctF1VztJ6shLfhg8wQCZpHwvlGMgHg8YNkM2ySiTYCpNOhGtN0nm0g6bx2P9foQT2b4swJFwQoHZ-Or1XOwuqKf0ROTXB0-XirHfoCcbfzrRJk_w_lCQyHyOzP2WMdUeXfVbfJis9Euw8oY3YrTjUe42Rt3FdWi388abbzEyOjWeGzGgsrjpwzX_zW8GQP0lx1fu1Q0NuK7nLMz9UrpzEdtbJmuZIOVmEvlbksZ2drsw2eB0x8KTKFFhg7XMGJkQQ1sZ9zO7oUTTC1paRbA' \
-d '{
"id": 123,
"hero": {
	"id": "string",
	"displayName": "string",
	"visibility": "string",
	"linkedAccounts": [
		{
			"provider": "string",
			"credential": "string"
		}
	],
	"profilePicture": {
		"url": "string",
		"alt": "string",
		"width": 0,
		"height": 0
	},
	"email": "string",
	"phoneNumber": "string",
	"name": {
		"first": "string",
		"second": "string",
		"last": "string"
	},
	"location": {
		"city": "string",
		"country": "string"
	},
	"userName": "string",
	"bio": "string",
	"workType": "string",
	"subWorkType": "string",
	"twitter": "string",
	"linkedin": "string",
	"website": "string",
	"bids": [
		{
			"id": "string",
			"bidderId": "string",
			"createdAt": "2023-10-02T21:29:12.443Z",
			"updatedAt": "2023-10-02T21:29:12.443Z",
			"rate": 0,
			"amount": 0,
			"currency": "string",
			"timeRequired": "string",
			"workingTime": "string",
			"questId": "string",
			"apprentice": {
				"rate": 0,
				"workingHours": {
					"start": 0,
					"end": 0,
					"timezone": "string"
				},
				"mentor": "string",
				"favoriteTo": "string",
				"workingDays": [
					"string"
				]
			},
			"apprenticeRate": 0,
			"apprenticeCut": 0,
			"totalEarnings": 0,
			"status": "string",
			"totalBidValue": 0
		}
	],
	"quests": [
		{
			"id": "string",
			"creatorId": "string",
			"reward": 0,
			"title": "string",
			"description": "string",
			"tags": [
				"TAG_DESIGN_UNSPECIFIED"
			],
			"images": [
				{
					"url": "string",
					"alt": "string",
					"width": 0,
					"height": 0
				}
			],
			"bidders": [
				"string"
			],
			"status": "string",
			"createdAt": "2023-10-02T21:29:12.443Z",
			"summary": "string",
			"level": 0
		}
	],
	"portfolios": [
		"string"
	],
	"experiences": [
		{
			"position": "string",
			"company": "string",
			"startDate": "string",
			"endDate": "string"
		}
	],
	"rating": 0,
	"isMentor": true,
	"mentor": {
		"bio": "string",
		"skill": "string",
		"maxRate": 0,
		"minRate": 0
	},
	"isApprentice": true,
	"apprentice": {
		"rate": 0,
		"workingHours": {
			"start": 0,
			"end": 0,
			"timezone": "string"
		},
		"mentor": "string",
		"favoriteTo": "string",
		"workingDays": [
			"string"
		]
	},
	"xp": "string",
	"region": "string",
	"language": "string",
	"level": 0,
	"isVerified": true,
	"twoFactorAuth": true,
	"receiveVerificationText": true,
	"notificationPreferences": {
		"questModified": true,
		"proposalReceived": true,
		"interviewAccepted": true,
		"interviewProposalDeclined": true,
		"offerAccepted": true,
		"questExpiresSoon": true,
		"questExpires": true,
		"noInterviews": true
	},
	"financials": {
		"balance": 0,
		"withdrawFrequency": "string",
		"withdrawThreshold": 0,
		"cardEnding": 0,
		"cardType": "string"
	}
}
}'

# curl -v -X 'GET' \
# 'http://localhost:4200/v1/heroes' \
# -H 'accept: application/json' \
# -H 'X-Firebase-ID-Token: eyJhbGciOiJSUzI1NiIsImtpZCI6IjlhNTE5MDc0NmU5M2JhZTI0OWIyYWE3YzJhYTRlMzA2M2UzNDFlYzciLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiUGlvdHIgT3N0cm93c2tpIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0xweF8yZ1dQMm0wVlByLVhiX2JyY1FxTEJpVGV2UDZyb1doWFM3Z2YzbHpUYz1zOTYtYyIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS9hc29sdGQtZ3VpbGRzIiwiYXVkIjoiYXNvbHRkLWd1aWxkcyIsImF1dGhfdGltZSI6MTY5NjI4MTM2NCwidXNlcl9pZCI6InpHZUs0aUxrRE1oU2t4cXR5eENOUFBJenNuNzMiLCJzdWIiOiJ6R2VLNGlMa0RNaFNreHF0eXhDTlBQSXpzbjczIiwiaWF0IjoxNjk2MjgxMzY0LCJleHAiOjE2OTYyODQ5NjQsImVtYWlsIjoicGlvdHJvc3RyQGdvb2dsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJnb29nbGUuY29tIjpbIjExNDc1ODU0OTgyNTE0NjI3MTQ5MyJdLCJlbWFpbCI6WyJwaW90cm9zdHJAZ29vZ2xlLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6Imdvb2dsZS5jb20ifX0.yy_J_9Jdiu7duO7fszlu1DLMdzim7C5qXW-Dkwzx1ZyjmlqKbWAdTOiA1L5T0HmttmIctF1VztJ6shLfhg8wQCZpHwvlGMgHg8YNkM2ySiTYCpNOhGtN0nm0g6bx2P9foQT2b4swJFwQoHZ-Or1XOwuqKf0ROTXB0-XirHfoCcbfzrRJk_w_lCQyHyOzP2WMdUeXfVbfJis9Euw8oY3YrTjUe42Rt3FdWi388abbzEyOjWeGzGgsrjpwzX_zW8GQP0lx1fu1Q0NuK7nLMz9UrpzEdtbJmuZIOVmEvlbksZ2drsw2eB0x8KTKFFhg7XMGJkQQ1sZ9zO7oUTTC1paRbA'
# works