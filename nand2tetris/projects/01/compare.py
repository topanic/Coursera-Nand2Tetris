from termcolor import colored

filenames = ['Not', 'And', 'Or', 'Xor', 'Not16', 'And16', 'Or16', 'Or8Way', 'DMux', 'DMux4Way', 'DMux8Way', 'Mux', 'Mux16', 'Mux4Way16', 'Mux8Way16']

# 对于文件中的每一行，打开两个文件
for filename in filenames:
    filename = filename.strip()
    try:
        with open(filename + '.out', 'r') as file1, open(filename + '.cmp', 'r') as file2:
            # 读取两个文件的内容
            content1 = file1.read()
            content2 = file2.read()

        # 比较两个文件的内容并以不同的颜色打印输出
        if content1 == content2:
            print(colored(f'{filename}: The files are identical.', 'green'))
        else:
            print(colored(f'{filename}: The files are different.', 'red'))
    except FileNotFoundError:
        print(colored(f'{filename}: The files are different.', 'red'))
    
    print('-----------------------------------')