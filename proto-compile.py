# windows 特供，用以编译 protobuf 的。
# 不能用通配符，太 TM 坑了
# 快被气死了
# 只能模拟一下啊，不然怎么办呢，操蛋！
import glob
from subprocess import call

def compile_protoc():
    proto_src_files = glob.glob(".\\proto-src\\*.proto")
    print("狗日的 Windows，连 TM 通配符都不能用，等等，我在编译...")

    protocs = ' '.join(proto_src_files)
    call("protoc -I proto-src/ {} --go_out=plugins=grpc:proto".format(protocs))

    print("老子终于编译完了！ 总共 {} 个 protobuf 文件".format(len(proto_src_files)))

compile_protoc()

