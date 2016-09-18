import os
import argparse
parser = argparse.ArgumentParser(description='genotype extract script for chr:pos format by Jacky')
parser.add_argument('-i','--id_list', help='file location for individial ID list',required=True)
parser.add_argument('-v','--snp_list',help='file location for SNP list in chr:pos format', required=True)
parser.add_argument('-j','--job_ID',help='ID for this session', required=True)
args = parser.parse_args()

mapping= open('mapping.f')
niagads={}
for indv in mapping:
	indv=indv.strip().split()
	niagads[indv[2]]=indv[0]+" "+indv[1]
mapping.close()

os.system("mkdir rsid"+args.job_ID)
ID=open(args.id_list)
ID_trim=open('./rsid'+args.job_ID+'/ID.trim','w')
for indv in ID:
	indv=indv.strip().split()
	if niagads.has_key(indv[0]):
		ID_trim.write('%s\n'%(niagads[indv[0]]))
ID.close()
ID_trim.close()

os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00017/ng00017 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00017")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00020/ng00020 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00020")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00022/ng00022 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00022")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00023/ng00023 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00023")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00024/ng00024 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00024")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00026/ng00026 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00026")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00028/ng00028 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00028")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00029/ng00029 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00029")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00030/ng00030 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00030")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00031/ng00031 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00031")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00032/ng00032 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00032")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00034/ng00034 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00034")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00042/a_ng00042 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.a_ng00042")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00042/b_ng00042 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.b_ng00042")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00042/c_ng00042 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.c_ng00042")
os.system("./plink -bfile /mnt/data/niagads_inquery/Inquery_NIAGADS_hg19_LQ/ng00043/ng00043 --keep ./rsid"+args.job_ID+"/ID.trim --extract "+args.snp_list+" --make-bed --out ./rsid"+args.job_ID+"/example.ng00043")
os.system("ls ./rsid"+args.job_ID+"/*bed | sed 's/.bed//g' >./rsid"+args.job_ID+"/merge-list")
os.system("./plink --merge-list ./rsid"+args.job_ID+"/merge-list  --recode transpose --out ./rsid"+args.job_ID+"/example.merged")

tfam=open('./rsid'+args.job_ID+'/example.merged.tfam')
ID_output=open('./rsid'+args.job_ID+'/ID.output','w')
ivd = dict((v, k) for k, v in niagads.items())
for indv in tfam:
        indv=indv.strip().split()
        if ivd.has_key(indv[0]+" "+indv[1]):
                ID_output.write('%s\t%s\t%s\n'%(indv[0],indv[1],ivd[indv[0]+" "+indv[1]]))
tfam.close()
ID_output.close()

succ_call={}
tped=open('./rsid'+args.job_ID+'/example.merged.tped')
for locus in tped:
        locus=locus.strip().split()
        succ_call[locus[1]]=1
tped.close()

pos=open(args.snp_list)
snp_na=open('./rsid'+args.job_ID+'/snp_na','w')
for locus in pos:
        locus=locus.strip().split()
        if locus[0] not in succ_call:
                snp_na.write('%s\n'%(locus[0]))
pos.close()
snp_na.close()
