// caid_demo.cc
#include<iostream> 
#include<caid_alg.h> 
int main(){
	caid::DeviceInfo device;
	device.boot_time_in_sec = "1600607106";
	device.country_code = "CN";
	device.language = "zh_hans_CN";


	  // 设备名称
      device.device_name = "aaaa";
      // 系统版本
      device.system_version = "13.1.1";
      // 设备machine
      device.machine = "iPhone10,3";
      // 运营商
      device.carrier_info = "中国移动";
      // 物理内存
      device.memory = "4047224832";
      // 磁盘容量
      device.disk = "127938088960";
      // 系统更新时间
      device.sys_file_time = "1595214620.383940";
      // 设备model
      device.model = "D22AP";
      // 时区
      device.time_zone = "28800";

	auto ret = caid::GenerateCAID(device);
	std::cout<<"Caid Size:"<<ret.size()<<std::endl;
}
