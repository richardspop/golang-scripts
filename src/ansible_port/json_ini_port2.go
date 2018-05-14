package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

var file_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\ansible_port\\cisco-vmr-cfg.json"
var output_loc = "C:\\Users\\porichar\\Work\\hackerrank\\src\\ansible_port\\tmp.ini"

var iniHeaderHost = "localhost   ansible_connection=local \n"
var iniHeaderTag = "[all]"
var iniHeaderTagVars = "[all:vars]"

// Static
var iniK8sFqdn = "vmr_k8s_master_fqdn=\"<filled with master fqdn>\""
var iniK8sPort = "vmr_k8s_master_port=\"<filled with master port>\""
var iniVmrVersion = "vmr_version_tag=\"<filled with VMR version>\""
var iniZKVersion = "vmr_zookeeper_tag=\"ZK_MHT_002\""

type platformTemplate struct {
	Path string `json:"path"`
	PackageName string `json:"packageName"`
}

type vmrTemplate struct {
	ReconPath string `json:"reconpath"`
	ArchivePath string `json:"archivepath"`
	ActivePath string `json:"activepath"`
	SiteId string `json:"siteId"`
	Locality string `json:"locality"`
	VSRM string `json:"vsrm"`
	K8sMaster string `json:"k8sMaster"`
	ObjstorePassword string `json:"objstorePassword"`
	ObjstoreUsername string `json:"objstoreUsername"`
	Objectstore []string `json:"objectstore"`
	ConsulIP string `json:"consulIP"`
	DbPassword string `json:"dbPassword"`
	DbUsername string `json:"dbUsername"`
	DatabasePort int `json:"databasePort"`
	DatabaseIP []string `json:"databaseIP"`
	DatabaseMasterPort int `json:"databaseMasterPort"`
	DatabaseMasterIP string `json:"databaseMasterIP"`
	DockerRegIp string `json:"dockerRegIp"`
	PlatformInfo platformTemplate `json:"platformInfo"`
	BulkDelete string `json:"bulkDelete"`
	ObjStoreVip string `json:"objStoreVip"`
	OptimizeObjStore string `json:"optimizeObjStore"`
	UsePutCopy bool `json:"usePutCopy"`
	AdaptationBaseURL bool `json:"adaptationBaseURL"`
	MaxChannels int `json:"maxChannels"`
	JitpServicePort int `json:"jitpServicePort"`
	RioApiPort int `json:"rioApiPort"`
	DataPlaneVip string `json:"dataPlaneVip"`
	HaproxyVip string `json:"haproxyVip"`
	VmrServiceIp string `json:"vmrServiceIp"`
	VmrServiceFqdn string `json:"vmrServiceFqdn"`
	OpenshiftRouterPort int `json:"openshiftRouterPort"`
	OpenshiftPlatform string `json:"openshiftPlatform"`
	JitpServiceHost string `json:"jitpServiceHost"`
	RioApiHost string `json:"rioApiHost"`
	SensuServerIP string `json:"sensuServerIP"`
	SensuServerPort string `json:"sensuServerPort"`
	SensuServerVHost string `json:"sensuServerVHost"`
	SensuServerUsername string `json:"sensuServerUsername"`
	SensuServerPassword string `json:"sensuServerPassword"`
	HttpsProxy string `json:"httpsProxy"`
	HttpProxy string `json:"httpProxy"`
	LogServerType string `json:"logServerType"`
	LogServer string `json:"logServer"`
	ElasticsearchPort int `json:"elasticsearchPort"`
	KafkaDefaultTopic string `json:"kafkaDefaultTopic"`
	KafkaIhPort string `json:"kafkaIhPort"`
}

func deleteFileIfExist(location string) (err error) {
	if _, err = os.Stat(location); err == nil {
		err = os.Remove(location)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeDataToINI(data *vmrTemplate, location string) error {
	w, err := os.Create(location)
	if err != nil {
		return err
	}
	defer w.Close()

	f := func(n int, localErr error) bool {
		err = localErr
		return err == nil
	}
	if !(f(fmt.Fprintln(w, iniHeaderHost)) &&
		f(fmt.Fprintln(w, iniHeaderTag)) &&
		f(fmt.Fprintln(w, iniHeaderTagVars)) &&
		f(fmt.Fprintln(w, iniK8sFqdn)) &&
		f(fmt.Fprintln(w, iniK8sPort))) {
		return err
	}

	if data.OpenshiftPlatform == "Yes" {
		if !f(fmt.Fprintln(w, "kubernetes_flavor=\"openshift\"")) {
			return err
		}
	} else {
		fmt.Println("Not supported")
		return fmt.Errorf("Only openshift is supported currectly")
	}

	if !(f(fmt.Fprintln(w, "vmr_docker_registry_ip=" + "\"" + data.DockerRegIp + "\"")) &&
		f(fmt.Fprintln(w, iniVmrVersion)) &&
		f(fmt.Fprintln(w, iniZKVersion)) &&
		f(fmt.Fprintln(w, "vmr_site_id=" + "\"" + data.SiteId + "\"")) &&
		f(fmt.Fprintln(w, "vmr_locality=" + "\"" + data.Locality + "\"")) &&
		f(fmt.Fprintln(w, "vmr_vsrm_endpoint=" + "\"" + strings.Split(data.VSRM, ":")[0] + "\"")) &&
		f(fmt.Fprintln(w, "vmr_vsrm_port="  + strings.Split(data.VSRM, ":")[1])) &&
		f(fmt.Fprintln(w, "vmr_k8s_master=" + "\"" + data.K8sMaster + "\"")) &&
		f(fmt.Fprintln(w, "vmr_service_ip=" + "\"" + data.VmrServiceIp + "\"")) &&
		f(fmt.Fprintln(w, "vmr_service_fqdn=" + "\"" + data.VmrServiceFqdn + "\"")) &&
		f(fmt.Fprintln(w, "vmr_oc_router_port=" + strconv.Itoa(data.OpenshiftRouterPort))) &&
		f(fmt.Fprintln(w, "vmr_jitp_service_host=" + "\"" + data.JitpServiceHost + "\"")) &&
		f(fmt.Fprintln(w, "vmr_api_host=" + "\"" + data.RioApiHost + "\"")) &&
		f(fmt.Fprintln(w, "vmr_max_channels=" + strconv.Itoa(data.MaxChannels))) &&
		f(fmt.Fprintln(w, "vmr_storage_endpoints=" + "\"" + strings.Join(data.Objectstore, ",") + "\"")) &&
		f(fmt.Fprintln(w, "vmr_storage_user=" + "\"" + data.ObjstoreUsername + "\"")) &&
		f(fmt.Fprintln(w, "vmr_storage_password=" + "\"" + data.ObjstorePassword + "\"")) &&
		f(fmt.Fprintln(w, "vmr_storage_active_path=" + "\"" + data.ActivePath + "\"")) &&
		f(fmt.Fprintln(w, "vmr_storage_archive_path=" + "\"" + data.ArchivePath + "\"")) &&
		f(fmt.Fprintln(w, "vmr_storage_recon_path=" + "\"" + data.ReconPath + "\"")) &&
		f(fmt.Fprintln(w, "vmr_storage_vip=" + "\"" + data.ObjStoreVip + "\"")) &&
		f(fmt.Fprintln(w, "vmr_memsql_master_endpoint=" + "\"" + data.DatabaseMasterIP + "\"")) &&
		f(fmt.Fprintln(w, "vmr_memsql_leaf_endpoint=" + "\"" + strings.Join(data.DatabaseIP, ",") + "\"")) &&
		f(fmt.Fprintln(w, "vmr_sensu_server_endpoint=" + "\"" + data.SensuServerIP + "\"")) &&
		f(fmt.Fprintln(w, "vmr_sensu_server_port=" + "\"" + data.SensuServerPort + "\"")) &&
		f(fmt.Fprintln(w, "vmr_sensu_server_vhost=" + "\"" + data.SensuServerVHost + "\"")) &&
		f(fmt.Fprintln(w, "vmr_sensu_server_user=" + "\"" + data.SensuServerUsername + "\"")) &&
		f(fmt.Fprintln(w, "vmr_sensu_server_password=" + "\"" + data.SensuServerPassword + "\"")) &&
		f(fmt.Fprintln(w, "vmr_log_server_type=" + "\"" + data.LogServerType + "\"")) &&
		f(fmt.Fprintln(w, "vmr_log_server_endpoint=" + "\"" + data.LogServer + "\"")) &&
		f(fmt.Fprintln(w, "vmr_elastic_search_port=" + strconv.Itoa(data.ElasticsearchPort))) &&
		f(fmt.Fprintln(w, "vmr_kafka_default_topic=" + "\"" + data.KafkaDefaultTopic + "\"")) &&
		f(fmt.Fprintln(w, "vmr_kafka_ih_endpoint=" + "\"" + data.KafkaIhPort + "\"")) &&
		f(fmt.Fprintln(w, "vmr_consul_endpoint=" + "\"" + data.ConsulIP + "\"")) &&
		f(fmt.Fprintln(w, "vmr_memsql_user=" + "\"" + data.DbUsername + "\"")) &&
		f(fmt.Fprintln(w, "vmr_memsql_password=" + "\"" + data.DbPassword + "\"")) &&
		f(fmt.Fprintln(w, "vmr_memsql_master_port=" + strconv.Itoa(data.DatabaseMasterPort))) &&
		f(fmt.Fprintln(w, "vmr_memsql_leaf_port=" + strconv.Itoa(data.DatabasePort))) &&
		f(fmt.Fprintln(w, "vmr_archive_agent_bulk_delete=" + "\"" + data.BulkDelete + "\"")) &&
		f(fmt.Fprintln(w, "vmr_storage_optimization_mode=" + "\"" + data.OptimizeObjStore + "\"")) &&
		f(fmt.Fprintln(w, "vmr_storage_use_put_copy=" + "\"" + strconv.FormatBool(data.UsePutCopy) + "\"")) &&
		f(fmt.Fprintln(w, "vmr_dash_origin_mpd_adaptation_base_url=" + "\"" + strconv.FormatBool(data.AdaptationBaseURL) + "\"")) &&
		f(fmt.Fprintln(w, "vmr_jitp_service_port=" + strconv.Itoa(data.JitpServicePort))) &&
		f(fmt.Fprintln(w, "vmr_api_port=" + strconv.Itoa(data.RioApiPort))) &&
		f(fmt.Fprintln(w, "vmr_haproxy_vip=" + "\"" + data.HaproxyVip + "\"")) &&
		f(fmt.Fprintln(w, "vmr_http_proxy=" + "\"" + data.HttpProxy + "\"")) &&
		f(fmt.Fprintln(w, "vmr_https_proxy=" + "\"" + data.HttpsProxy + "\""))) {
		return err
	}


	return nil
}

func main() {
	f, err := ioutil.ReadFile(file_loc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	err = deleteFileIfExist(output_loc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	data := new(vmrTemplate)
	json.Unmarshal(f, &data)

	fmt.Println(data.OpenshiftPlatform)
	fmt.Println(writeDataToINI(data, output_loc))
}
