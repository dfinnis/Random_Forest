import matplotlib.pyplot as plt
import csv

# read_csv reads accuracy.csv
def read_csv():
	train_accuracy = []
	test_accuracy = []
	with open('accuracy.csv', 'r') as file:
		reader = csv.reader(file)
		for row in reader:
			train_accuracy.append(float(row[0]))
			test_accuracy.append(float(row[1]))
	return train_accuracy, test_accuracy

# visualize plots accuracy over training
def visualize(train_accuracy, test_accuracy):
	depth = list(range(1, len(test_accuracy) + 1))

	plt.plot(depth, train_accuracy, label='train accuracy')
	plt.plot(depth, test_accuracy, label='test accuracy')
	plt.title('Accuracy over Depth')
	plt.xlabel('Depth')
	plt.ylabel('Accuracy')
	plt.legend()
	plt.show()

# main reads accuracy.csv & plots accuracy
def main():
	try:
		train_accuracy, test_accuracy = read_csv()
		visualize(train_accuracy, test_accuracy)
	except Exception:
		print("Error: Failed to visualize data. Is data valid?")
		pass

if __name__ == '__main__':
	main()
