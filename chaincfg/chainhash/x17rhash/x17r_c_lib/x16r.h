#ifndef __X16R_H__
#define __X16R_H__
#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

#include <stddef.h>

void x16r_hash(void* output, void* input, const int in_len);
#ifdef __cplusplus
}
#endif

#endif /* __X16R_H__ */
