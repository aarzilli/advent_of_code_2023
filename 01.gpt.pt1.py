import sys                                                                                                                                               
import re                                                                                                                                                
                                                                                                                                                          
def calibrate(filename):                                                                                                                                 
     with open(filename, 'r') as f:                                                                                                                       
         lines = f.readlines()                                                                                                                            
     total_calibration = 0                                                                                                                                
     for line in lines:                                                                                                                                   
         numbers = re.findall(r'\d', line)                                                                                                                
         if len(numbers) > 1:                                                                                                                             
             calibration = int(numbers[0] + numbers[-1])                                                                                                  
         elif len(numbers) == 1:                                                                                                                          
             calibration = int(numbers[0]) * 11  # If there's only one number, consider that as both first and last digit                                 
         else:                                                                                                                                            
             calibration = 0                                                                                                                              
         total_calibration += calibration                                                                                                                 
     return total_calibration                                                                                                                             
                                                                                                                                                          
if __name__ == "__main__":                                                                                                                               
     filename = sys.argv[1]                                                                                                                               
     print(calibrate(filename))                                                                                                                           