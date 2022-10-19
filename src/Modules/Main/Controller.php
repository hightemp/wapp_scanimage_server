<?php

namespace Hightemp\WappSnanimageServer\Modules\Main;

use Hightemp\WappSnanimageServer\Project;

class Controller
{
    public static function fnIndexHTML()
    {
        $sScannersList = Project::fnGetScannersListCached();
        $aScannedFiles = Project::fnGetScannedFiles();
        $aPackedFiles = Project::fnGetArchivedFiles();

        require_once("view/index.php");
    }

    public static function fnImageHTML()
    {
        if (isset($_GET['file'])) {
            $sPath = Project::fnGetImageRelPath($_GET['file']);
            require_once("view/image_view.php");
        } else {
            require_once("view/blank.php");
        }
    }

    public static function fnBlankHTML()
    {
        require_once("view/blank.php");
    }

    public static function fnLoaderHTML()
    {
        require_once("view/loader.php");
    }

    public static function fnActionHTML()
    {
        if (isset($_GET['action'])) {
            if ($_GET['action']=='scan_image') {
                $sImageFile = Project::fnScanImage();
                echo <<<HTML
<script>
window.parent.frames[1].location = '?m=Main&a=fnImageHTML&file={$sImageFile}';
window.parent.location.reload();
</script>
HTML;
            }
        }
        if (isset($_POST['action'])) {
            if ($_POST['action']=='scan') {
                echo <<<HTML
<script>
window.parent.frames[1].location = '?m=Main&a=fnLoaderHTML';
window.location = '?m=Main&a=fnActionHTML&action=scan_image';
</script>
HTML;
            }
            if ($_POST['action']=='delete') {
                foreach ($_POST['images'] as $sFile) {
                    $sFilePath = Project::fnGetImagePath($sFile);
                    unlink($sFilePath);
                }
                echo <<<HTML
<script>window.parent.location.reload()</script>
HTML;
            }
            if ($_POST['action']=='delete_all') {
                $aFiles = Project::fnGetScannedFiles();
                foreach ($aFiles as $aFile) {
                    unlink($aFile[1]);
                }
                echo <<<HTML
<script>window.parent.location.reload()</script>
HTML;
            }

            if ($_POST['action']=='delete_archives') {
                foreach ($_POST['archives'] as $sFile) {
                    $sFilePath = Project::fnGetArchivePath($sFile);
                    unlink($sFilePath);
                }
                echo <<<HTML
<script>window.parent.location.reload()</script>
HTML;
            }
            if ($_POST['action']=='delete_all_archives') {
                $aFiles = Project::fnGetArchivedFiles();
                foreach ($aFiles as $aFile) {
                    unlink($aFile[1]);
                }
                echo <<<HTML
<script>window.parent.location.reload()</script>
HTML;
            }

            if ($_POST['action']=='download') {
                $zip = new \ZipArchive;
                if ($zip->open(Project::C_ARCHIVED_PATH.'/'.time().'.zip', \ZipArchive::CREATE) === TRUE)
                {
                    foreach ($_POST['images'] as $sFile) {
                        $sFilePath = Project::fnGetImagePath($sFile);
                        $zip->addFile($sFilePath);
                    }
                    
                    $zip->close();
                }
                echo <<<HTML
<script>window.parent.location.reload()</script>
HTML;
            }
            if ($_POST['action']=='download_all') {
                $zip = new \ZipArchive;
                if ($zip->open(Project::C_ARCHIVED_PATH.'/'.time().'.zip', \ZipArchive::CREATE) === TRUE)
                {
                    $aFiles = Project::fnGetScannedFiles();
                    foreach ($aFiles as $aFile) {
                        $sFilePath = Project::fnGetImagePath($aFile[1]);
                        $zip->addFile($sFilePath);
                    }
                    
                    $zip->close();
                }
                echo <<<HTML
<script>window.parent.location.reload()</script>
HTML;
            }
            
        }

        die();
    }
}