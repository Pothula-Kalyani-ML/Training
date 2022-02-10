from unittest import result
import matrix_read
list_of_matrices=matrix_read.read_matrix_from_file("f1")

mat1=list_of_matrices[0]
mat2=list_of_matrices[1]
print(mat1)
print(mat2)
if (len(mat1)==len(mat2)) and (len(mat1[0])==len(mat2[0])):
  result_mat=[[ 0 for x in mat1],[0 for x in mat1[0]]]
  print(result_mat)
  for i in range(len(mat1)):  
      for j in range(len(mat1[0])):
        result_mat[i][j] = mat1[i][j] + mat2[i][j]
  print(result_mat)
else:
    print("addition not possible")
