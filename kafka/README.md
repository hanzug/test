考虑一个电子商务网站，需要处理用户的订单、支付信息、产品库存等数据。这里，我们将演示如何使用Kafka来构建一个简化的订单处理系统。

场景描述
订单生成器（Order Generator）：模拟用户下单行为，将订单信息发送到Kafka主题中。

库存服务（Inventory Service）：监听订单主题，处理订单并更新产品库存。

支付服务（Payment Service）：监听订单主题，处理订单支付信息。

日志服务（Logging Service）：监听订单主题，记录订单处理日志。

Kafka主题

使用三个Kafka主题：

orders：用于接收订单信息。

inventory_updates：用于发送库存更新信息。

payment_updates：用于发送支付信息更新。