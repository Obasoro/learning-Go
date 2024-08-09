# Define the data to store employee's data

hourly_wage_employees = []
monthly_wage_employees= []

# Function for Hourly rate employees
def cal_hourly_earnings(hourly_rate, hours_worked):
    return hourly_rate * hours_worked

# Function for monthly net wage

def cal_monthly_net(monthly_wage):
    tax_rate = 0.10
    insurance_deduction = 100
    return monthly_wage - (tax_rate*monthly_wage) - (insurance_deductio)


# Function to calculate the employee data using hourly rate

def hourly_rate_data():
    employee_name = input("Enter employee's data name: ")
    hourly_rate = float(input("Enter the hourly rate: "))
    hours_worked = float(input("Enter the number of hours worked: "))
    total_earning = cal_hourly_earnings(hourly_rate, hours_worked)
    hourly_wage_employees.append((employee_name, total_earning))

# Function for monthly wage data

def monthly_rate_data():
    employee_name = input("Enter the employee's data: ")
    monthly_wage = float(input("Enter the monthly wage: "))
    net_pay = cal_monthly_net(monthly_wage)
    monthly_wage_enployees.append((employee_name, monthly_wage, net_pay))


# Data informatiom for hourly employee
num_hourly_employees = int(input("Enter the number of hourly employees: "))
for _, in range(num_hourly_employees):
    hourly_rate_data()


# Data information for monthly employee

num_monthly_employees = int(input("Enter the number of monthly employee: "))
for _, in range (num_monthly_employees):
    monthly_rate_data()

# Generate the report
print("\nEmployee Report:")
print("#####################")
print("#####################")
print("Hourly Employees:")
for name, earnings in hourly_wage_employees:
    print(f"{name}: Earnings - ${earnings:.2f}")

print("\nSalaried Employees:")
for name, salary, net_pay in monthly_wage_employees:
    print(f"{name}: Monthly Salary - ${salary:.2f}, Net Pay - ${net_pay:.2f}")


