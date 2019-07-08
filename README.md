# kafkalib

This library contains mostly functions for mapping textual representations (tab delimited/json) of
kafka types to the sarama kafka library. It is intended to help making cli applications that output info
about kafka primitives or parse json representations of them.

- (un)marshaling of json to flat or nested intermediary structs
- marshaling of intermediary structs to text
- (un)marshaling of the intermediary structs to sarama types