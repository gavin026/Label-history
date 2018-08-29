#coding=utf-8
import sys,os,subprocess,shutil,re
import xlwt, xlrd
from xlutils.copy import copy
from xlwt import Style
from config import *
import time
rootpath = sys.path[0]
TemplateFile  = os.path.join(rootpath, "import-template-syzkaller.xls")
sheetName = "Sheet1"
SEC_Fuzz = "SEC_Fuzz_Secu_Spec_0303_"

def get_raw_data(excelFile, sheetName):
	rbook = xlrd.open_workbook(excelFile,formatting_info=True)
	rsh = rbook.sheet_by_name(sheetName)
	if rsh.cell_value(0, 56).strip() != "TestCase_name":
		print "wrong excel file!"
		sys.exit(1)
	wbook = copy(rbook)
	w_sheet = wbook.get_sheet(0)
	for row in range(1, rsh.nrows):
		TestCase_name = rsh.cell_value(row, 56).strip()
		os.system("sed -i " + "'/" + TestCase_name + "/d' add_cbglist.txt") #删除add_cbglist中的TestCase_name
		w_sheet.write(row,0,"")   
		w_sheet.write(row,56,"")   
		w_sheet.write(row,57,"")   
		w_sheet.write(row,58,"")   
		w_sheet.write(row,59,"")   
	wbook.save(TemplateFile)
	NUM = int(rsh.cell_value(row, 57).strip().split('_')[5])
	return NUM
		
if __name__ == "__main__":
	start = time.time()
	os.system("sh run.sh")
	Index = get_raw_data(TemplateFile,sheetName)
	rbook = xlrd.open_workbook(TemplateFile,formatting_info=True)
	wbook = copy(rbook)
	w_sheet = wbook.get_sheet(0)
	row = 0
	file = open("add_cbglist.txt")
	for line in file.readlines():
		line=line.strip("\n")
		row += 1
		Index += 1
		print line
		w_sheet.write(row,56,line)
		w_sheet.write(row,97,TestStep1 + line +".cfg\n" + TestStep2)
		w_sheet.write(row,57,SEC_Fuzz + str(Index))
		w_sheet.write(row,0,Arr[0])
		w_sheet.write(row,58,Arr[1])
		w_sheet.write(row,59,Arr[2])
	wbook.save(TemplateFile)
	end = time.time()
	print "Successful!"
	print "程序运行耗时：" + str(end-start) + "s"

	
