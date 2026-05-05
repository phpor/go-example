#ifndef ATTRIBUTION_CAID_ALGORITHM_CAID_ALG_H_
#define ATTRIBUTION_CAID_ALGORITHM_CAID_ALG_H_

#include <string>
#include <utility>
#include <vector>

namespace caid {

struct DeviceInfo {
  // 设备启动时间
  std::string boot_time_in_sec;
  // 国家
  std::string country_code;
  // 语言
  std::string language;
  // 设备名称
  std::string device_name;
  // 系统版本
  std::string system_version;
  // 设备machine
  std::string machine;
  // 运营商
  std::string carrier_info;
  // 物理内存
  std::string memory;
  // 磁盘容量
  std::string disk;
  // 系统更新时间
  std::string sys_file_time;
  // 设备model
  std::string model;
  // 时区
  std::string time_zone;
};

// brief:
// @input : 设备指纹
// @return: pair<version, caid>形式的多版本caid
//          Example: {{"20200901", "f949f306494646edfee0f939698e1fb8"},
//                    {"20200909", "d7019d00248fe6729426804fdf93c4fa"}}
//          其中version以算法版本发布日期形式提供
std::vector<std::pair<std::string, std::string> > GenerateCAID(const DeviceInfo& device);

}  // namespace caid

#endif  // ATTRIBUTION_CAID_ALGORITHM_CAID_ALG_H_
