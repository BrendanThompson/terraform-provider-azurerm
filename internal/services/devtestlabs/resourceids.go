package devtestlabs

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Schedule -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.DevTestLab/schedules/schedule1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DevTestLab -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DevTestLabPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/policySets/policyset1/policies/policy1
