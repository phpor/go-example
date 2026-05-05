#include <iostream>

#include "caid_alg.h"

int main() {
  // Step 1: 构建device
  caid::DeviceInfo device;
  // 设备启动时间
  device.boot_time_in_sec = "1600607106";
  // 国家
  device.country_code = "CN";
  // 语言
  device.language = "zh-Hans-CN";
  // 设备名称
  device.device_name = "E910DDDB2748C36B47FCDE5DD720EEC1";
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
  // Step 2: 计算caid
  //         结果类型: std::vector<std::pair<std::string, std::string>>
  //         这里将返回一个caid结果: f949f306494646edfee0f939698e1fb8
  //         对应版本为: 20200901
  //         后续算法升级后将返回多个版本的caid
  auto ret = caid::GenerateCAID(device);
  std::cout << "Caid Size:" << ret.size() << std::endl;
  for (auto& caid : ret) {
    // caid类型: std::pair<std::string, std::string>
    //           first是版本号, second是caid结果
    std::cout << "Caid version:" << caid.first << "|value:" << caid.second << std::endl;
  }
  return 0;
}
