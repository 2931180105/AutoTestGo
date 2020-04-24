# 生成abi文件、再编译出bytecode，然后添加到生成的go文件中，最后直接调用部署函数
abigen --abi=Store_sol_Store.abi --pkg=store --out=Store.go >Store_sol_Store.abi
solc --bin Store.sol >Store_sol_Store.bin
abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go