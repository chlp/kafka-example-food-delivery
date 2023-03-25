package refrigerator

// Idea is not to link business logic module (refrigerator) with the kafka module,
// so we don't get a constant from kafka.
// Additionally, GotNewFoodHandler can receive data not only from kafka.
const ctxMsgKey = "msg-id" //

var Names = []string{"Salad", "Sandwich", "Bread", "Steak", "Tuna Steak", "Fish", "Shrimp", "Rice", "Spaghetti", "Pizza", "Hamburger", "Eggs", "Cheese", "Sausages", "Apple Juice", "Grape Juice", "Milk", "Candy", "Cookie", "Pie", "Cake", "Cupcake"}
