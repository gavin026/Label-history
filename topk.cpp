#include <iostream>
#include <vector>

using namespace std;

class MinHeap {
	
private:
	int maxsize;                //堆的大小
	void filterDown(int begin); //向下调整堆
	vector<int> arr;

public:
	MinHeap(int k);          //构造函数
	void createMinHeap(int arr[]); //创建堆
	void insert(int val);    //插入元素
	int getTop();            //获取堆顶元素
	vector<int> getHeap();  //获取堆中的全部元素
};

MinHeap::MinHeap(int k) {
	maxsize = k;
}

void MinHeap::createMinHeap(int a[]) {
	for(int i = 0; i < maxsize; i++) {
		arr.push_back(a[i]);
	}
	
	for(int i = arr.size()/2-1;i>=0;i--) {
		filterDown(i);
	}
}

void MinHeap::insert(int val) {
	if(val > getTop()) {
		arr[0] = val;
		filterDown(0);
	}
}

void MinHeap::filterDown(int current) {
	int end = arr.size() - 1;
	int child = current * 2 + 1;  //当前节点的左孩子
	int val = arr[current];       //保存当前节点
	while(child <= end) {
		//选出两个孩子中较小的孩子
		if(child < end && arr[child+1] < arr[child])
			child++;
		if(val < arr[child]) break;
		else {
			arr[current] = arr[child]; //孩子节点覆盖当前节点
			current = child;
			child = child *2 + 1;
		}
	}
	arr[current] = val;
}

int MinHeap::getTop() {
	if(arr.size() != 0)
		return arr[0];
	return NULL;
}

vector<int> MinHeap::getHeap() {
	vector<int> heap;
	for(int i = 0;i<arr.size();i++)
		heap.push_back(arr[i]);
	return heap;
}


int main() {
    // Test case
    int arr[] = {7,6,1,2,9,8,5,12,23,44,21,33,0,4,3};
    int k = 4;
    MinHeap heap(4); // 创建一个大小为4的堆 
    heap.createMinHeap(arr);
	
    for(int i = k; i < 15; i++) {
        heap.insert(arr[i]);
    }

    cout << "最大的四个元素是" << endl;
    vector<int> v = heap.getHeap();
    for(int i = 0; i < v.size(); i++) {
        cout << v[i] << endl;
    }
    return 0;
}
