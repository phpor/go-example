%module caid
%include "stl.i"
%{
#include "caid_alg.h"
%}
namespace std {
  %template(CaidVector) std::vector<std::pair<std::string, std::string>>;
  %template(CaidMap) std::pair<std::string, std::string>;
}
%include "caid_alg.h"
