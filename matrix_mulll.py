def matrix_input():
    input_file=open("mat-input","a")
    r=int(input( "enter the number of rows:"))
    c=int(input("enter the number of column:"))
    matrix=[]
    
    for i in range(r):
        a=[]
        for j in range(c):
            a.append(int(input()))
        matrix.append(a) 
    for i in matrix :
           input_file.writelines('%s\n' % i)
    
    input_file.close()
    return matrix
def mat_multiplication(mat1,mat2):
    if len(mat1[0])!=len(mat2):
        print("multiplication not possible")
        exit(1)
    result_mat=[]
    for i in range(len(mat1)):
      row=[]
      for j in range(len(mat2[0])):
         row.append(0)
      result_mat.append(row)      
    for i in range(len(mat1)):
        for j in range(len(mat2[0])):
            for k in range(len(mat2)):
                result_mat[i][j] += mat1[i][k] * mat2[k][j]
    return result_mat
if __name__ == "__main__":
    matrix1=matrix_input()
    matrix2=matrix_input()
    result_matrix=mat_multiplication(matrix1,matrix2) 
    print(result_matrix)


