// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package sophos

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "sophos", asset.ModuleFieldsPri, AssetSophos); err != nil {
		panic(err)
	}
}

// AssetSophos returns asset data.
// This is the base64 encoded gzipped contents of module/sophos.
func AssetSophos() string {
	return "eJzUXMFy4zYSvc9XoJLLpCrj3dmjD1vl2J4dV40nWklOjiwIbJFYgwADgJLl0/zDXpOfmy/ZAkjKpAQKtNRKZj0nUZrXj92N7kY3yHfkETaXxKgyV+YNIZZbAZfku/rCd28IScEwzUvLlbwk/3xDCGl+Te5VWgl4Q8iSg0jNpf/uHZG0gA6i+7ObEi5JplVVNlcCqH2cLtZTtr0UwnJ/GgRQA5dkAZZ2rqewpJWwiYe+JEsqDPS+DtCo/+q7I0ulSUm14TJr7ukpI2ZjhMouOr/fpd6ln8KKM+h91d7GI2zWSqc73x1gVX/dwwvLpDYsMfBFRNwNtUDebjabzbuieJemP5B1DpLYHAisQFqiGKu0hvQwI8sLeFYSUQ9zXgB5m+eXRXFpzJG0amUm7gMes3uVgiCyKhagiVp6UqOsVpPhu+JOoDIDzekxXITKUIk8SP5bBeT9PwjLqabMgjaEqRTI27+/d//e/xAn5OQjOtCmBKeS2l/gIrsgS65hTYWor8UJMVWUSnZ/ejKr6xaSaDClkoYvmjgkVJa5OOSZfmiZ6joIR4iaaoGrvFm18P99q8DDHPJFkgMVNsdj8BGotguglhhLbWUOEyg1V5rbDebSWoFDJAJWIPza0nS55OwwkR2yJ9N4EJYXLkbXwB0e5OuX/5IrIdQaUqI0uQHJo/Gw0tRJCvITSmavTB6VpgsunJY6vN4aYEqmJrLil+vEOfdQGOLSQgb6dYS2y2bqsvvdDVnnnOWEG0LLUnCnqDqJjLJlZUAj5459yGHJu1XQiaL/5eA8PrGqUYxThZNFFuCsH1lknCLSuZMWtARLrhgDY0ipBGcbZ7NjTMVLk9QIqP50N5mdSKwhhRudJzUlH6BbUlaNJ0XLcsmFBX0OnV05QsyHGfLBS2lVeIz+6Ata0uP8f0AYz+BdivHo0dUZbvg6nofm5hGPx5SbxyYxU2N4Jl8WQEj/UXYWWC6VUBliATHfYraF+VHUGLWQKY1IrGtEblwdqsQKUrLYEKdKaisNrqAwG8lyrSR/hvRV3FvExTejzYSbhAlVIe54drR43UMP5yiZuPiklxSzX3DXQvpNBJdMFW4X0cSqH/124kcyUdqSq8P8VGXPT1BVNlNDBH+K1NaaJTxcguxdjpD6WfOMSyqIUZVmQO4mhKapdoXI6Epfs6SgDE9Ru5zur66PIsVUJa3eJG77jblxTaFdeo2Efh35oslRxWRqLLoxUzCWy3pNHmFRR+kvUt4O81EaLLWyiimBWFc2iN2G0livK5UOt0mOqsl214JDd3Tm1xNCZUoebibjbXomal2bHc2Ps6JE3h3cXd9Pts2b+sNoJrhu74WzxvfHMzEgbVI+2nAb5fVtirmyvSZpSdkjWOPlHGaigQFfQfpnsGlljdDNYmPhXHQ89jjVrP4EIuO0YjWVJnFhaD95HJdT5g5RUAtpoEQIFTJjCeIGo32aPhIdSxAxKXeYDaTlU0ieTYt7Qf31LI1mz0rCeaYZjZWdgLZ0/nT1eRQhPDKfabFLJpqFz6eSrsm6evk1ppeGFb5edhlFaHCdpNwg9lbnE59THDIw96MLMlGmHjWtqKjAXH798rvS2dcvf/xIvn75XUMpNu2Hr1/+OEyYKSn7Q6GTGd/WQ125LYm5IUJlbmedgQTt1macFGofsxlp8hSk5UteJyV/6yzec1ih07neSiZ3N+1eoqDGgh7Niqfl+bvkJ41cOgxxO5cdjn4Ewi/gYu/iSdS3vTPMwfps25C7uxkrvjC7ldcpY9AtgQKMMTSLBDMmqDF8id4Bf+ExICFcpFYCEvzZ8KRB9NuarQ8dplIKapdKF4gsGsQ2GDReehExEHo/2WngpXHc4kdoWKozQEwgxynDZzHMczBX84nPZWCjCxbKpKowZd/KtFRcWvLwEJNtcw3UIs+G5pMQbjhY+cLx2+nldgq2b6wnKVTGZVKZvZx8oqU8LunhDq8Q3Lrdia8PZfVww+KfgFV+DlZSzNNFnkOLTXrYQR50dZ6y5J6KNdVADKNSuh0mWjmy1KpIoKBcJI3PISZjkClo4tHJLno44KhzcZmCCyTg4t4r+Jhq8R9gmBsYL3sXNijb/dLw57AXHVWA7yGG9yRcY55M83AjTgF8mE+SSiNOKT7MJ+Rh+sl7eLNGai5rakiq1lIomsa2i47VdnuMx+2mhXQB3hH1Lasl6EvyUDpe/qRewzGygLkAXDeZ8eftFMqhE5tT6zaPlnIJ6Y6DDJLCjcQTavMeqYaPC4YjCeGG5A+OhI+/1BjFuG8LrrnNX06fRyjZkqmioBKxtnOu1IC6pJ3Wx+Ff3H6pKhnxJ9Q1ePL6S1VBucS13I3HPI3XbxXVVFqO2Rb0Lu4s1zprz99fJMZGP5ol51Bbk81r6BEBPTX2LDym9ajnFUw0UIMZvacej6zzjTePBqZ0WrsPWGAuEFBDTEmLvxVUcMZVLDxpWILGrN2nO4BhTylpQZEz26ykBbka0WR0xQjmtnYHL7JqE2yf+PcWetfdwqr3x/SR5+iz+ux/DzQoXa1Aa55CYtUjICrh5waX9HGDHJj6hvrwW4XQyuZK82fMtbjVSgA8vAtyxWC9MBFXyPwFtTMauPrFbyojDbDKF6WJS0TIp1CacreuM6P7/C4P3KzS5RHPKLm1ZfNgFgyu4defbWjPwHyczyftg1+jHgH5k5//GNk7PtOTei08qXxBUlNcw8JwC2RJhYgdKWxqvHO4c7tnG+vQ8MTAg+HtuD9xY7eHCHNgj8ZJEVVan91ew6IjNWZJJS16Y6+1Y33K0QsYGZ1xl/w2Lo94RIFZvuJ2g0vgV1hsHyFpBNR77YJalrsqUqaEUb+Rc8pqfqvBVMJGLNd3cuwNVNfJR7Txc/reVIgjrdnHq/eNb1fbUQ63UJAFcJkRKqnYPMc2SgX4QSXqA873NWZ0quMKE0y5vbH/qNtGv+dIoPNt3TP1U2dj+ql+IeA2pXwPKD4VSI2t96qIS7AzLrrpYw/fPG6b0N/8LNpO7hSzqCutX84e5tAUB6iRuwc4rHP8pmM83vqk+mSTUsOSP6EGGZesyaSPG4753MIZHkqbuRrvehf2oBZMtcTXwpMlsz5ueOVz76BUYw8Eb7bA/jRhvHrB7HXUp/DGNHlwB9IP0TF0a3NfQOGb/L4HO2Dx8w2LIlaubJ4wwVEPXF5VNifXfdBh6QWwnEpuEEs9T2Afd7CoqjWFu9iuw7hBDhoKZUGCXSuN+CRzGDbIQChGRUYtrCli4P0UQh2Wj66ATyHUiBfgbluvw7hBDkqkK9AGNQz8vI8ZlC1hjS778z5mUDYv0Q9u3O1Bhq3uY1RS5hvDGT3DAZI6CJJWwLjzIw2rXBmLXAA0dBzymAkYXScp7b0U70T5U7omPcSg3HvU+cZ9dLBhwBjcrjmZ7UGGJVuqreUDJn79C/dme3gH0g3Ww1hTj0buJoeFOlrG0iIs9PX3uo8XFDvbmE8qS2a3019up8nnq/tbRCt7bFJjkx52kMuCsseqTApU//7Jg5Ii7ub+BCvi3ffxBpaWXmEW87M+XlCmi254Ej920QZWUz1yGVzHrx/pTEOQA3WLeuSIRt3BC8r8rQK9MVbzvds6QXAINCgdnqxGTEl9uKDEycfJ7HY26/SLTpa6DzmcHxLsBEHiblWfzMaTW2/+x2SmFLXfN729ifb5FppKliNXWj950BFtlqp02kyRcvE+WrjiTgUkrKyCMpdC0d0AGhHr8Mj330eceWMsFJhia8SIfg3ov0Am4hmVHlpYnuSIGa+HFi7llKUiKaAY6tIe96C8f0fDfR81KH+pAXEk0UML36+m0hTcWkhBa6URN4gdaLKDPRAg61dWYPNocV9F4nHB0d5csiXQB43ZA5NC1xavY5FqVZ6DRR/3oC0wKWxtMUI+U0JwM3gW5ZjjTHuA4RyG/za5fcig5Gsllzw79LbjI/JKGDQon0yhVHrA7Y8QvYcXlLp9SBpL7D5gUO4cBloGR4jsYQ2WC4hhtQ8XLosMZsHbQwvKw3zPMo1os+5kmsS/sAL3RhvoTuO+vvf/BQAA//8gM69j"
}
