/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/A-Zee029/rocketmq-client-go/v2/admin"
	"github.com/A-Zee029/rocketmq-client-go/v2/primitive"
	"strings"
	"time"
)

func main() {
	var topic, nameServerAddrs, clusterName, brokerAddr string

	flag.StringVar(&topic, "topic", "newTopic", "topic to operate")
	flag.StringVar(&nameServerAddrs, "nameSrvAddrs", "127.0.0.1:9876", "nameServerAddress")
	flag.StringVar(&clusterName, "clusterName", "defaultCluster", "clusterName")
	flag.StringVar(&brokerAddr, "broker", "127.0.0.1:10911", "brokerAddress")

	flag.Parse()

	nameSrvAddr := strings.Split(nameServerAddrs, ";")

	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	if err != nil {
		fmt.Println(err.Error())
	}

	//create topic
	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		//admin.WithBrokerAddrCreate(brokerAddr),
		admin.WithClusterNameCreate(clusterName),
	)
	if err != nil {
		fmt.Println("Create topic error:", err.Error())
	}

	time.Sleep(5 * time.Second)
	topicList, err := testAdmin.FetchAllTopicList(context.Background())
	if err != nil {
		fmt.Println("List topic error:", err.Error())
	}
	fmt.Println("topicList:")
	for _, topicName := range topicList.TopicNameList {
		fmt.Println(topicName)
	}

	//deletetopic
	err = testAdmin.DeleteTopic(
		context.Background(),
		admin.WithTopicDelete(topic),
		admin.WithClusterNameDelete(clusterName),
	)
	if err != nil {
		fmt.Println("Delete topic error:", err.Error())
	}

	topicList, err = testAdmin.FetchAllTopicList(context.Background())
	if err != nil {
		fmt.Println("List topic error:", err.Error())
	}
	fmt.Println("topicList:")
	for _, topicName := range topicList.TopicNameList {
		fmt.Println(topicName)
	}

	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}

}
