import sys                                                                                                                                               
import re                                                                                                                                                
                                                                                                                                                          
def calibrate(filename):                                                                                                                                 
     numbers = {                                                                                                                                          
         'zero': '0',                                                                                                                                     
         'one': '1',                                                                                                                                      
         'two': '2',                                                                                                                                      
         'three': '3',                                                                                                                                    
         'four': '4',                                                                                                                                     
         'five': '5',                                                                                                                                     
         'six': '6',                                                                                                                                      
         'seven': '7',                                                                                                                                    
         'eight': '8',                                                                                                                                    
         'nine': '9'                                                                                                                                      
     }                                                                                                                                                    
     total_calibration = 0                                                                                                                                
     with open(filename, 'r') as f:                                                                                                                       
         lines = f.readlines()                                                                                                                            
     for line in lines:                                                                                                                                   
         for word, digit in numbers.items():                                                                                                              
             print(line, word, digit)
             line = line.replace(word, digit)                                                                                                             
         print(line)
         digits = re.findall(r'\d', line)                                                                                                                 
         if digits:                                                                                                                                       
             print(int(digits[0] + digits[-1]))
             total_calibration += int(digits[0]+digits[-1])                                                                                               
     return total_calibration                                                                                                                             
                                                                                                                                                          
if __name__ == "__main__":                                                                                                                               
     filename = sys.argv[1]                                                                                                                               
     print(calibrate(filename))                                                                                                                           
                                                                                                                                                          