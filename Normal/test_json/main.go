package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
)

type json_get struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				ContainerLabelComDockerComposeConfigHash         string `json:"container_label_com_docker_compose_config_hash"`
				ContainerLabelComDockerComposeContainerNumber    string `json:"container_label_com_docker_compose_container_number"`
				ContainerLabelComDockerComposeOneoff             string `json:"container_label_com_docker_compose_oneoff"`
				ContainerLabelComDockerComposeProject            string `json:"container_label_com_docker_compose_project"`
				ContainerLabelComDockerComposeProjectConfigFiles string `json:"container_label_com_docker_compose_project_config_files"`
				ContainerLabelComDockerComposeProjectWorkingDir  string `json:"container_label_com_docker_compose_project_working_dir"`
				ContainerLabelComDockerComposeService            string `json:"container_label_com_docker_compose_service"`
				ContainerLabelComDockerComposeVersion            string `json:"container_label_com_docker_compose_version"`
				ContainerLabelRestartcount                       string `json:"container_label_restartcount"`
				ID                                               string `json:"id"`
				Image                                            string `json:"image"`
				Instance                                         string `json:"instance"`
				Job                                              string `json:"job"`
				Name                                             string `json:"name"`
			} `json:"metric"`
			Values [][]interface{} `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

func main() {

	starttime := "2023-08-29T02:40:18.000Z"
	//fmt.Println("")
	endtime := "2023-08-29T02:41:33.000Z"
	step := "1s"

	//  可更改

	url := "http://43.138.142.78/prometheus/api/v1/query_range?query=sum%28irate%28container_cpu_usage_seconds_total%7Bimage=%22jaegertracing/jaeger-collector%22%7D[5m]%29%29%20without%20%28cpu%29&start="
	url += starttime
	url += "&end="
	url += endtime
	url += "&step="
	url += step

	// 发送GET请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("GET请求失败:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	file, err := os.Create(fmt.Sprintf("%s\n", starttime))

	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	// 打印响应内容
	var json_ json_get
	err = json.Unmarshal(body, &json_)
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(fmt.Sprintf("URL: %s\n", url))
	file.WriteString(fmt.Sprintf("开始时间: %s\n", starttime))
	file.WriteString(fmt.Sprintf("结束时间: %s\n", endtime))
	file.WriteString(fmt.Sprintf("步长: %s\n", step))
	Write_file(json_, file)

	file.Close()
}

func Write_file(json_ json_get, file *os.File) {
	var get_95 []float64
	var sum float64
	min := math.Inf(1)
	max := math.Inf(-1)
	for _, row := range json_.Data.Result[0].Values {
		file.WriteString(fmt.Sprintf("时间戳: %.0f\n", row[0]))
		file.WriteString(fmt.Sprintf("值: %s\n", row[1]))

		value := row[1].(string)
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum += f
		get_95 = append(get_95, f)

		if f < min {
			min = f
		}

		if f > max {
			max = f
		}
	}

	file.WriteString(fmt.Sprintf("\n平均值: %f\n", sum/float64(len(json_.Data.Result[0].Values))))
	file.WriteString(fmt.Sprintf("最大值: %f\n", max))
	file.WriteString(fmt.Sprintf("最小值: %f\n", min))

	sort.Float64s(get_95)
	index := 0.95 * float64(len(get_95))
	fmt.Println(get_95)
	fmt.Println(index)
	file.WriteString(fmt.Sprintf("95CPU: %f\n", get_95[int(index)]))

}
